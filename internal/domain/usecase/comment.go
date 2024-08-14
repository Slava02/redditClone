package usecase

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/pkg/hexid"
	"redditClone/pkg/logger"
)

type CommentUseCase struct {
	service interfaces.ICommentService
}

var _ interfaces.ICommentUseCase = &CommentUseCase{}

func NewCommentUseCase(service interfaces.ICommentService) *CommentUseCase {
	return &CommentUseCase{
		service: service,
	}
}

func (c CommentUseCase) AddComment(ctx context.Context, postID string, comment entities.Comment) (entities.PostExtend, error) {
	const op = "internal.usecases.comment.AddComment: "

	id, err := hexid.Generate()
	if err != nil {
		logger.Error(op + "couldn't generate id")

		return entities.PostExtend{}, fmt.Errorf("%w", IdGenerateError)
	}

	commentExt := entities.NewCommentExtend(comment, id)

	post, err := c.service.AddComment(ctx, postID, commentExt)
	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return post, nil
}

func (c CommentUseCase) DeleteComment(ctx context.Context, username string, postID string, commentID string) (entities.PostExtend, error) {
	const op = "internal.usecase.comment.DeleteComment: "

	post, err := c.service.DeleteComment(ctx, username, postID, commentID)
	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%s: %w", op, err)
	}

	return post, nil
}
