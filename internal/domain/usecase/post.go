package usecase

import (
	"context"
	"fmt"
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
	const op = "internal.usecase.GetPosts"

	posts, err := p.service.GetPosts(ctx)
	if err != nil {
		return []entities.PostExtend{}, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

func (p PostUseCase) AddPost(ctx context.Context, post entities.Post) (entities.PostExtend, error) {

	panic("implement me")
}

func (p PostUseCase) GetPostsWithCategory(ctx context.Context, category string) ([]entities.PostExtend, error) {
	const op = "internal.usecase.GetPostsWithCategory"

	posts, err := p.service.GetPostsWithCategory(ctx, category)
	if err != nil {
		return []entities.PostExtend{}, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

func (p PostUseCase) GetPostsWithUser(ctx context.Context, username string) ([]entities.PostExtend, error) {

	panic("implement me")
}

func (p PostUseCase) GetPost(ctx context.Context, postID string) (entities.PostExtend, error) {
	const op = "internal.usecase.GetPost"

	post, err := p.service.GetPost(ctx, postID)
	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return post, nil
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
