package inMemory

import (
	"context"
	"errors"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/pkg/logger"
	"sync"
)

var ErrExists = errors.New("already exists")

type Users struct {
	counterID int
	mutex     sync.RWMutex
	users     map[string]entities.UserExtend
}

var _ interfaces.IUserRepository = &Users{}

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

	if _, exists := u.users[user.ID]; !exists {
		u.users[user.ID] = user
	} else {
		logger.Errorf(op + ErrExists.Error())

		return fmt.Errorf("%w", ErrExists)
	}
	u.counterID++

	return nil
}

func (u Users) Get(ctx context.Context, username string) (entities.UserExtend, error) {

	panic("implement me")
}

func (u Users) Contains(ctx context.Context, username string) (bool, error) {

	panic("implement me")
}
