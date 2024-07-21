package inMemory

import (
	"context"
	"os/user"
	"redditClone/internal/domain/entities"
)

// реализуем интерфейс storage

type Posts struct{}

func NewPosts() *Posts {
	return &Posts{}
}

func (p Posts) Posts(ctx context.Context) ([]*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p Posts) PostsByCategory(ctx context.Context, category string) ([]*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p Posts) PostsByUser(ctx context.Context, user user.User) ([]*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p Posts) PostById(ctx context.Context, postID string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p Posts) CreatePost(ctx context.Context, item entities.Post, author user.User) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p Posts) DeletePost(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}

func (p Posts) UpVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p Posts) DownVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p Posts) UnVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}
