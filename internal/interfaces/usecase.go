package interfaces

import (
	"context"
	"redditClone/internal/domain/entities"
)

type IUserUseCase interface {
	SignUp(ctx context.Context, user entities.User) (entities.UserExtend, error)
	Login(ctx context.Context, username string, password string) (entities.UserExtend, error)
}

type IPostUseCase interface {
	GetPosts(ctx context.Context) ([]entities.PostExtend, error)
	AddPost(ctx context.Context, post entities.Post) (entities.PostExtend, error)
	GetPostsWithCategory(ctx context.Context, category string) ([]entities.PostExtend, error)
	GetPostsWithUser(ctx context.Context, username string) ([]entities.PostExtend, error)
	GetPost(ctx context.Context, postID string) (entities.PostExtend, error)
	DeletePost(ctx context.Context, username string, postID string) error
	Upvote(ctx context.Context, userID string, postID string) (entities.PostExtend, error)
	Downvote(ctx context.Context, userID string, postID string) (entities.PostExtend, error)
	Unvote(ctx context.Context, userID string, postID string) (entities.PostExtend, error)
}

type ICommentUseCase interface {
	AddComment(ctx context.Context, postID string, comment entities.Comment) (entities.PostExtend, error)
	DeleteComment(ctx context.Context, username string, postID string, commentID string) (entities.PostExtend, error)
}
