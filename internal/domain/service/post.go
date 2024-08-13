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
	const op = "internal.service.AddPost: "

	err := p.repo.Add(ctx, post)
	if err != nil {
		logger.Errorf(op, err.Error())

		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p PostService) GetPost(ctx context.Context, postID string) (entities.PostExtend, error) {
	const op = "internal.service.GetPost: "

	post, err := p.repo.Get(ctx, postID)
	if err != nil {
		logger.Errorf(op, err.Error())

		return entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return post, nil
}

func (p PostService) GetPosts(ctx context.Context) ([]entities.PostExtend, error) {
	const op = "internal.service.GetPosts: "

	posts, err := p.repo.GetAll(ctx)
	if err != nil {
		logger.Errorf(op, err.Error())

		return []entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return posts, nil
}

func (p PostService) GetPostsWithCategory(ctx context.Context, category string) ([]entities.PostExtend, error) {
	const op = "internal.service.GetPostsWithCategory: "

	posts, err := p.repo.GetWhereCategory(ctx, category)
	if err != nil {
		logger.Errorf(op, err.Error())

		return []entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return posts, nil
}

func (p PostService) GetPostsWithUser(ctx context.Context, username string) ([]entities.PostExtend, error) {
	const op = "internal.service.GetPostsWithUser: "

	posts, err := p.repo.GetWhereUsername(ctx, username)
	if err != nil {
		logger.Errorf(op, err.Error())

		return []entities.PostExtend{}, fmt.Errorf("%w", err)
	}

	return posts, nil
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
	const op = "internal.service.DeletePost: "

	post, err := p.repo.Get(ctx, postID)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if post.Author.Username != username {
		return fmt.Errorf("%w", ErrNotAllowed)
	}

	err = p.repo.Delete(ctx, postID)
	if err != nil {
		logger.Errorf(op, err.Error())

		return fmt.Errorf("%w", err)
	}

	return nil
}

func (p PostService) SortPostsByTime(posts []entities.PostExtend) []entities.PostExtend {

	panic("implement me")
}
