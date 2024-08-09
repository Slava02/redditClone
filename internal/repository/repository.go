package repository

import (
	"errors"
	"redditClone/internal/interfaces"
)

var (
	ErrExists   = errors.New("already exists")
	ErrNotFound = errors.New("not found")
)

type Repositories struct {
	PostRepository interfaces.IPostRepository
	UserRepository interfaces.IUserRepository
}
