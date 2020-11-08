package main

import (
	cache "chaelub/thinksurance/small-project/app/cache/user"
	repoConfig "chaelub/thinksurance/small-project/app/repo"
	"chaelub/thinksurance/small-project/app/repo/role"
	repo "chaelub/thinksurance/small-project/app/repo/user"
	"chaelub/thinksurance/small-project/app/server"
	store "chaelub/thinksurance/small-project/app/store/user"
	"database/sql"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Config struct {
	Migrations struct {
		Dir string `toml:"dir"`
	} `toml:"migrations"`
	HttpServer    server.HTTPServerConfig `toml:"httpInterface"`
	DBStoreConfig repoConfig.DBRepoConfig `toml:"dbStore"`
}

var (
	configPath *string
)

func init() {
	configPath = flag.String("config", "./config/dev.toml", "path to toml config file")
	flag.Parse()
}

func main() {
	conf := Config{}

	// Parse config
	_, err := toml.DecodeFile(*configPath, &conf)
	if err != nil {
		log.Fatalf("can't parse config file [%s]: %v", *configPath, err)
	}

	// init DB connection
	db, err := sql.Open(conf.DBStoreConfig.Driver, conf.DBStoreConfig.DSN)
	if err != nil {
		log.Fatalf(
			"can't open connection to [%s] with driver [%s]: %v",
			conf.DBStoreConfig.Driver,
			conf.DBStoreConfig.DSN,
			err)
	}
	defer db.Close()

	// migrate DB
	err = initDB(db, conf.Migrations.Dir)
	if err != nil {
		panic(err)
	}

	// init user store with user repository and user cache
	userRepo := repo.NewDBUserRepo(db)
	userCache := cache.NewInMemoryUserCache()
	userStore := store.NewUserStoreWIthCache(userCache, userRepo)

	// it's possible to create a role store with cache. Like the userStore
	// but avoid it to simplify the project
	roleRepo := role.NewDBRoleRepo(db)

	// run http server as a service interface
	go server.RunServer(conf.HttpServer, userStore, roleRepo)

	// need to stop execution. In real word should listen OS signals
	select {}
}

func initDB(db *sql.DB, migrationsDir string) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsDir,
		"postgres", driver)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err == os.ErrNotExist {
		version, _, err := m.Version()
		if err != nil {
			return fmt.Errorf("cannot get current database version")
		}

		return fmt.Errorf("No migration file for current database version. Current version %d", version)
	}

	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	version, _, err := m.Version()
	if err != nil {
		return fmt.Errorf("cannot get current database version")
	}

	log.Printf("Database version: %d\n", version)

	return nil
}
