package repository

import (
	"errors"
	"redditClone/internal/interfaces"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrExists         = errors.New("exists")
	ErrBadCredentials = errors.New("invalid login or password")
)

type Repositories struct {
	PostRepository interfaces.IPostRepository
	UserRepository interfaces.IUserRepository
}
