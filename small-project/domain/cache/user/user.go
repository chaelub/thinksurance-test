package user

import "chaelub/thinksurance/small-project/domain/models"

type UserCacheI interface {
	Get(id int64) (models.User, error)
	Save(id int64, user models.User) error
	Delete(id int64) error
}
