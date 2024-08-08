package usecase

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
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

	panic("implement me")
}

func (c CommentUseCase) DeleteComment(ctx context.Context, username string, postID string, commentID string) (entities.PostExtend, error) {

	panic("implement me")
}
