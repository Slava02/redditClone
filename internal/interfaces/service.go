package interfaces

import (
	"context"
	"redditClone/internal/domain/entities"
)

type IUserService interface {
	AddUser(ctx context.Context, user entities.UserExtend) error
	GetUser(ctx context.Context, username string) (entities.UserExtend, error)
	ContainsUser(ctx context.Context, username string) (bool, error)
}

type IPostService interface {
	AddPost(ctx context.Context, post entities.PostExtend) error
	GetPost(ctx context.Context, postID string) (entities.PostExtend, error)
	GetPosts(ctx context.Context) ([]entities.PostExtend, error)
	GetPostsWithCategory(ctx context.Context, category string) ([]entities.PostExtend, error)
	GetPostsWithUser(ctx context.Context, username string) ([]entities.PostExtend, error)
	UpvotePost(ctx context.Context, userID string, postID string) (entities.PostExtend, error)
	DownvotePost(ctx context.Context, userID string, postID string) (entities.PostExtend, error)
	UnvotePost(ctx context.Context, userID string, postID string) (entities.PostExtend, error)
	DeletePost(ctx context.Context, username string, postID string) error
	SortPostsByTime(posts []entities.PostExtend) []entities.PostExtend
}

type ICommentService interface {
	AddComment(ctx context.Context, postID string, comment entities.CommentExtend) (entities.PostExtend, error)
	DeleteComment(ctx context.Context, username string, postID string, commentID string) (entities.PostExtend, error)
}
