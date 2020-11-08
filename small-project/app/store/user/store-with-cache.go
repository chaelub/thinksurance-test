package user

import (
	cache "chaelub/thinksurance/small-project/app/cache/user"
	userI "chaelub/thinksurance/small-project/domain/cache/user"
	"chaelub/thinksurance/small-project/domain/models"
	repoI "chaelub/thinksurance/small-project/domain/repo/user"
	storeI "chaelub/thinksurance/small-project/domain/store/user"
	"errors"
)

type userStoreWithCache struct {
	cache userI.UserCacheI
	repo  repoI.UserRepoI
}

func (s *userStoreWithCache) GetByLogin(login string) (models.User, error) {
	user, err := s.repo.GetUserByLogin(login)
	if err != nil {
		return user, err
	}
	return user, s.cache.Save(user.Id, user)
}

func (s *userStoreWithCache) GetById(id int64) (models.User, error) {
	user, err := s.cache.Get(id)
	if err != nil && errors.Is(err, cache.ErrNotFound) {
		return user, err
	} else if errors.Is(err, cache.ErrNotFound) {
		user, err = s.repo.GetUserById(id)
		if err != nil {
			return user, err
		}
		s.cache.Save(user.Id, user)
	}
	return user, nil
}

func NewUserStoreWIthCache(cache userI.UserCacheI, repo repoI.UserRepoI) storeI.UserStoreI {
	return &userStoreWithCache{
		cache: cache,
		repo:  repo,
	}
}
