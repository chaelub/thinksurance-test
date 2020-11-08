package user

import (
	"chaelub/thinksurance/small-project/domain/models"
	"chaelub/thinksurance/small-project/domain/repo/user"
	"database/sql"
)

type dbUserRepo struct {
	db *sql.DB
}

func (s *dbUserRepo) GetUserByLogin(login string) (models.User, error) {
	row := s.db.QueryRow(userByLoginWithRole, login)

	var u models.User
	return u, row.Scan(&u.Id, &u.Login, &u.Password, &u.RoleId)
}

func (s *dbUserRepo) GetUserById(id int64) (models.User, error) {
	row := s.db.QueryRow(userByIDWithRole, id)

	var u models.User
	return u, row.Scan(&u.Id, &u.Login, &u.Password, &u.RoleId)
}

func NewDBUserRepo(db *sql.DB) user.UserRepoI {
	return &dbUserRepo{
		db: db,
	}
}
