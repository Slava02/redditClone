package usecase

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
)

type PostUseCase struct {
	service interfaces.IPostService
}

var _ interfaces.IPostUseCase = &PostUseCase{}

func NewPostUseCase(service interfaces.IPostService) *PostUseCase {
	return &PostUseCase{
		service: service,
	}
}

func (p PostUseCase) GetPosts(ctx context.Context) ([]entities.PostExtend, error) {

	panic("implement me")
}

func (p PostUseCase) AddPost(ctx context.Context, post entities.Post) (entities.PostExtend, error) {

	panic("implement me")
}

func (p PostUseCase) GetPostsWithCategory(ctx context.Context, category string) ([]entities.PostExtend, error) {

	panic("implement me")
}

func (p PostUseCase) GetPostsWithUser(ctx context.Context, username string) ([]entities.PostExtend, error) {

	panic("implement me")
}

func (p PostUseCase) GetPost(ctx context.Context, postID string) (entities.PostExtend, error) {

	panic("implement me")
}

func (p PostUseCase) DeletePost(ctx context.Context, username string, postID string) error {

	panic("implement me")
}

func (p PostUseCase) Upvote(ctx context.Context, userID string, postID string) (entities.PostExtend, error) {

	panic("implement me")
}

func (p PostUseCase) Downvote(ctx context.Context, userID string, postID string) (entities.PostExtend, error) {

	panic("implement me")
}

func (p PostUseCase) Unvote(ctx context.Context, userID string, postID string) (entities.PostExtend, error) {

	panic("implement me")
}
