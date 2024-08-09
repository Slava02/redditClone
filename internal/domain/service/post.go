package service

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/pkg/logger"
)

type PostService struct {
	repo interfaces.IPostRepository
}

func NewPostService(repo interfaces.IPostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (p PostService) AddPost(ctx context.Context, post entities.PostExtend) error {
	panic("implement me")
}

func (p PostService) GetPost(ctx context.Context, postID string) (entities.PostExtend, error) {
	panic("implement me")
}

func (p PostService) GetPosts(ctx context.Context) ([]entities.PostExtend, error) {
	const op = "service.GetPosts: "

	posts, err := p.repo.GetAll(ctx)
	if err != nil {
		logger.Errorf(op, err.Error())

		return []entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return posts, nil
}

func (p PostService) GetPostsWithCategory(ctx context.Context, category string) ([]entities.PostExtend, error) {
	const op = "service.GetPostsWithCategory: "

	posts, err := p.repo.GetWhereCategory(ctx, category)
	if err != nil {
		logger.Errorf(op, err.Error())

		return []entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return posts, nil
}

func (p PostService) GetPostsWithUser(ctx context.Context, username string) ([]entities.PostExtend, error) {

	panic("implement me")
}

func (p PostService) UpvotePost(ctx context.Context, userID string, postID string) (entities.PostExtend, error) {

	panic("implement me")
}

func (p PostService) DownvotePost(ctx context.Context, userID string, postID string) (entities.PostExtend, error) {

	panic("implement me")
}

func (p PostService) UnvotePost(ctx context.Context, userID string, postID string) (entities.PostExtend, error) {

	panic("implement me")
}

func (p PostService) DeletePost(ctx context.Context, username string, postID string) error {

	panic("implement me")
}

func (p PostService) SortPostsByTime(posts []entities.PostExtend) []entities.PostExtend {

	panic("implement me")
}
