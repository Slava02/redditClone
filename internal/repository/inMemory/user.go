package inMemory

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/domain/service"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
	"time"
)

type Users struct {
	counterID int
	users     map[string]*entities.User
}

var _ service.UserRepository = &Users{}

func NewUsers() *Users {
	users := map[string]*entities.User{
		"test": {
			ID:           1,
			Username:     "test",
			Password:     "test1234",
			RegisteredAt: time.Now(),
		},
	}
	return &Users{users: users}
}

func (u *Users) Get(ctx context.Context, userName, passwordHash string) (*entities.User, error) {
	user, ok := u.users[userName]

	logger.Infof("Getting user: %s Pass: %s", userName, passwordHash)
	//  TODO: не секьюр, надо как-то отрефакторить
	if !ok || passwordHash != user.Password {
		return nil, repository.ErrBadCredentials
	}

	return user, nil
}

func (u *Users) UserExists(ctx context.Context, userName string) bool {
	_, exists := u.users[userName]

	return exists
}

func (u *Users) SetSession(ctx context.Context, userID int, session entities.Session) error {
	//  TODO: реализовать stateful хранение сессий и подумать в каком сервисе это сделать
	return nil
}

func (u *Users) NextID(ctx context.Context) int {
	return u.counterID + 1
}

func (u *Users) AddUser(ctx context.Context, user *entities.User) error {
	u.users[user.Username] = user

	u.counterID++

	return nil
}
