package usecase

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/pkg/hexid"
	"redditClone/pkg/logger"
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
	const op = "internal.usecase.AddPost"

	id, err := hexid.Generate()
	if err != nil {
		logger.Error(op + "couldn't generate id")

		return entities.PostExtend{}, fmt.Errorf("%w", IdGenerateError)
	}

	postExtend := entities.NewPostExtend(post, id)

	err = p.service.AddPost(ctx, postExtend)

	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return postExtend, nil
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
	const op = "internal.usecase.GetPostsWithUser"

	posts, err := p.service.GetPostsWithUser(ctx, username)
	if err != nil {
		return []entities.PostExtend{}, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
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
	const op = "internal.usecase.DeletePost"

	err := p.service.DeletePost(ctx, username, postID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
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
