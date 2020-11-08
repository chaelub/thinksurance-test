package user

import "chaelub/thinksurance/small-project/domain/models"

type UserRepoI interface {
	GetUserByLogin(login string) (models.User, error)
	GetUserById(id int64) (models.User, error)
	// Expand interface with methods like SaveUser(models.User) and so on
}
