package repo

type DBRepoConfig struct {
	Driver string `toml:"driver"`
	DSN    string `toml:"dsn"`
}
