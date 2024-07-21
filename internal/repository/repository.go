package repository

import (
	"redditClone/internal/domain/service"
	"redditClone/internal/repository/inMemory"
)

const inmemory = 1

type Repositories struct {
	CommentRepository service.CommentRepository
	PostRepository    service.PostRepository
	UserRepository    service.UserRepository
}

func NewRepositories(t int) *Repositories {
	switch t {
	case inmemory:
		return &Repositories{
			CommentRepository: inMemory.NewComments(),
			PostRepository:    inMemory.NewPosts(),
			UserRepository:    inMemory.NewUsers(),
		}
	default:
		return nil
	}
}
