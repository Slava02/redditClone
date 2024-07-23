package inMemory

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/domain/service"
)

type Comments struct{}

var _ service.CommentRepository = &Comments{}

func NewComments() *Comments {
	return &Comments{}
}

func (c *Comments) CreateComment(ctx context.Context, comment string, itemID string, author entities.User) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Comments) DeleteComment(ctx context.Context, commentID string, itemID string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}
