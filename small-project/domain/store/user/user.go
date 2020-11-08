package user

import "chaelub/thinksurance/small-project/domain/models"

type UserStoreI interface {
	GetByLogin(login string) (models.User, error)
	// to get user from store after login we can use method GetByID for example
}
