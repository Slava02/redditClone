package service

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/pkg/logger"
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
	const op = "internal.service.comment.AddComment"

	post, err := c.repo.AddComment(ctx, postID, comment)
	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return post, nil
}

func (c CommentService) DeleteComment(ctx context.Context, username string, postID string, commentID string) (entities.PostExtend, error) {
	const op = "internal.service.DeleteComment: "

	post, err := c.repo.GetComment(ctx, postID, commentID)
	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	if post.Author.Username != username {
		return entities.PostExtend{}, fmt.Errorf("%w", ErrNotAllowed)
	}

	postExt, err := c.repo.DeleteComment(ctx, postID, commentID)
	if err != nil {
		logger.Errorf(op, err.Error())

		return entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return postExt, nil
}
