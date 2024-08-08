package service

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
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

	panic("implement me")
}

func (p PostService) GetPostsWithCategory(ctx context.Context, category string) ([]entities.PostExtend, error) {

	panic("implement me")
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
