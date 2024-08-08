package service

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
)

type CommentService struct {
	repo interfaces.IPostRepository
}

func NewCommentService(repo interfaces.IPostRepository) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (c CommentService) AddComment(ctx context.Context, postID string, comment entities.CommentExtend) (entities.PostExtend, error) {
	panic("implement me")
}

func (c CommentService) DeleteComment(ctx context.Context, username string, postID string, commentID string) (entities.PostExtend, error) {
	panic("implement me")
}
