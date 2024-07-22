package inMemory

import (
	"context"
	"redditClone/internal/domain/entities"
)

type Users struct {
	counterID int
	users     map[string]*entities.User
}

func NewUsers() *Users {
	return &Users{}
}

func (u *Users) CheckUserExists(ctx context.Context, userName string) bool {
	_, res := u.users[userName]

	return res
}

func (u *Users) SetSession(ctx context.Context, userID int, session entities.Session) error {

}

func (u *Users) NextID(ctx context.Context) int {
	return u.counterID + 1
}

func (u *Users) AddUser(ctx context.Context, user *entities.User) error {
	u.users[user.Username] = user

	u.counterID++

	return nil
}
