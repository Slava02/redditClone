package inMemory

import (
	"context"
	"os/user"
	"redditClone/internal/domain/entities"
)

type Comments struct{}

func NewComments() *Comments {
	return &Comments{}
}

func (c *Comments) CreateComment(ctx context.Context, comment string, itemID string, author user.User) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Comments) DeleteComment(ctx context.Context, commentID string, itemID string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}
