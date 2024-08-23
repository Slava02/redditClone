package inMemory

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
	"sync"
)

type Users struct {
	counterID int
	mutex     sync.RWMutex
	users     map[string]entities.UserExtend
}

var _ interfaces.IUserRepository = (*Users)(nil)

func NewUsers() *Users {
	users := map[string]entities.UserExtend{
		"test1234": {
			ID: "userid775slava",
			User: entities.User{
				Username: "test1234",
				Password: "test1234",
			},
		},
	}
	return &Users{
		counterID: 2,
		users:     users,
	}
}

func (u Users) Add(ctx context.Context, user entities.UserExtend) error {
	const op = "repository.inmemory.Add: "

	u.mutex.RLock()
	defer u.mutex.RUnlock()

	if _, exists := u.users[user.Username]; !exists {
		u.users[user.Username] = user
	} else {
		logger.Info(op, fmt.Sprintf("user %s already exists", user.Username))

		return fmt.Errorf("%w", repository.ErrExists)
	}
	u.counterID++

	return nil
}

func (u Users) Get(ctx context.Context, username string) (entities.UserExtend, error) {
	const op = "repository.inmemory.Get: "

	u.mutex.RLock()
	defer u.mutex.RUnlock()

	if _, exists := u.users[username]; !exists {
		logger.Info(op, fmt.Sprintf("user %s not found", username))

		return entities.UserExtend{}, fmt.Errorf("%w", repository.ErrNotFound)
	}

	return u.users[username], nil
}

func (u Users) Contains(ctx context.Context, username string) (bool, error) {

	panic("implement me")
}
