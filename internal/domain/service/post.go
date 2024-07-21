package service

import (
	"context"
	"os/user"
	"redditClone/internal/domain/entities"
)

type PostRepository interface {
	Posts(ctx context.Context) ([]*entities.Post, error)
	PostsByCategory(ctx context.Context, category string) ([]*entities.Post, error)
	PostsByUser(ctx context.Context, user user.User) ([]*entities.Post, error)
	PostById(ctx context.Context, postID string) (*entities.Post, error)
	CreatePost(ctx context.Context, item entities.Post, author user.User) (*entities.Post, error) // Здесь используется только category, text, title, type, нужно DTO
	DeletePost(ctx context.Context, ID string) error
	UpVotePost(ctx context.Context, id string) (*entities.Post, error)
	DownVotePost(ctx context.Context, id string) (*entities.Post, error)
	UnVotePost(ctx context.Context, id string) (*entities.Post, error)
}

type PostService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (s *PostService) Posts(ctx context.Context) ([]*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) PostsByCategory(ctx context.Context, category string) ([]*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) PostsByUser(ctx context.Context, user user.User) ([]*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) PostById(ctx context.Context, postID string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) CreatePost(ctx context.Context, item entities.Post, author user.User) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) DeletePost(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) UpVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) DownVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *PostService) UnVote(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}
