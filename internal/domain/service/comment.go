package service

import (
	"context"
	"redditClone/internal/domain/entities"
)

type CommentRepository interface {
	CreateComment(ctx context.Context, comment string, itemID string, author entities.User) (*entities.Post, error)
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
