package service

import (
	"context"
	"os/user"
	"redditClone/internal/domain/entities"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment string, itemID string, author user.User) (*entities.Post, error)
	DeleteComment(ctx context.Context, commentID string, itemID string) (*entities.Post, error)
}

type CommentService struct {
	repo CommentRepository
}

func NewCommentService(repo CommentRepository) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) CreateComment(ctx context.Context, comment string, itemID string, author user.User) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (s *CommentService) DeleteComment(ctx context.Context, commentID string, itemID string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}
