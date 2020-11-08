package user

import (
	"chaelub/thinksurance/small-project/domain/cache/user"
	"chaelub/thinksurance/small-project/domain/models"
	"errors"
	"sync"
)

var (
	ErrNotFound = errors.New("not found")
)

type inMemoryUserCache struct {
	mu    sync.RWMutex
	users map[int64]models.User
}

func (i *inMemoryUserCache) Get(id int64) (models.User, error) {
	i.mu.RLock()
	user, got := i.users[id]
	i.mu.RUnlock()
	if !got {
		return user, ErrNotFound
	}
	return user, nil
}

func (i *inMemoryUserCache) Save(id int64, user models.User) error {
	i.mu.Lock()
	i.users[id] = user
	i.mu.Unlock()
	return nil
}

func (i *inMemoryUserCache) Delete(id int64) error {
	i.mu.Lock()
	delete(i.users, id)
	i.mu.Unlock()
	return nil
}

func NewInMemoryUserCache() user.UserCacheI {
	return &inMemoryUserCache{
		users: make(map[int64]models.User),
	}
}
