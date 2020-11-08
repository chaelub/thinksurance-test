package role

import (
	"chaelub/thinksurance/small-project/domain/models"
	"chaelub/thinksurance/small-project/domain/repo/role"
	"database/sql"
	"fmt"
)

type dbRoleRepo struct {
	db *sql.DB
}

func (s *dbRoleRepo) GetRoleById(id int64) (models.Role, error) {
	rows, err := s.db.Query(roleWithPermissions, id)
	if err != nil {
		return models.Role{}, err
	}
	defer rows.Close()

	role := models.Role{}
	for rows.Next() {
		perm := models.Permission{}
		rows.Scan(&role.Id, &role.Name, &perm.Id, &perm.Source, &perm.Mode)
		role.Permissions = append(role.Permissions, perm)
	}
	fmt.Println(role)
	return role, rows.Err()
}

func NewDBRoleRepo(db *sql.DB) role.RoleRepoI {
	// todo: check migrations
	return &dbRoleRepo{
		db: db,
	}
}
