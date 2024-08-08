package inMemory

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
)

type Users struct {
	counterID int
	users     map[string]*entities.UserExtend
}

var _ interfaces.IUserRepository = &Users{}

func NewUsers() *Users {
	users := map[string]*entities.UserExtend{
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

	panic("implement me")
}

func (u Users) Get(ctx context.Context, username string) (entities.UserExtend, error) {

	panic("implement me")
}

func (u Users) Contains(ctx context.Context, username string) (bool, error) {

	panic("implement me")
}
