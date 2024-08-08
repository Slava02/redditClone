package interfaces

import (
	"context"
	"redditClone/internal/domain/entities"
)

type IUserRepository interface {
	Add(ctx context.Context, user entities.UserExtend) error
	Get(ctx context.Context, username string) (entities.UserExtend, error)
	Contains(ctx context.Context, username string) (bool, error)
}

type IPostRepository interface {
	Add(ctx context.Context, post entities.PostExtend) error
	Get(ctx context.Context, postID string) (entities.PostExtend, error)
	GetWhereCategory(ctx context.Context, category string) ([]entities.PostExtend, error)
	GetWhereUsername(ctx context.Context, username string) ([]entities.PostExtend, error)
	GetAll(ctx context.Context) ([]entities.PostExtend, error)
	Update(ctx context.Context, postID string, newPost entities.PostExtend) error
	Delete(ctx context.Context, postID string) error

	AddComment(ctx context.Context, postID string, comment entities.CommentExtend) (entities.PostExtend, error)
	GetComment(ctx context.Context, postID string, commentID string) (entities.CommentExtend, error)
	DeleteComment(ctx context.Context, postID string, commentID string) (entities.PostExtend, error)
}
