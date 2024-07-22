package repository

import (
	"errors"
	"redditClone/internal/domain/service"
)

var (
	ErrNotFound = errors.New("not found")
	ErrExists   = errors.New("exists")
)

type Repositories struct {
	CommentRepository service.CommentRepository
	PostRepository    service.PostRepository
	UserRepository    service.UserRepository
}
