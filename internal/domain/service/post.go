package service

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/pkg/logger"
)

type PostRepository interface {
	GetAllPosts(ctx context.Context) ([]*entities.Post, error)
	GetPostsByCategory(ctx context.Context, category string) ([]*entities.Post, error)
	GetCategories(ctx context.Context) ([]string, error)
	PostsByUser(ctx context.Context, user entities.User) ([]*entities.Post, error)
	PostById(ctx context.Context, postID string) (*entities.Post, error)
	CreatePost(ctx context.Context, item entities.Post, author entities.Author) (*entities.Post, error) // Здесь используется только category, text, title, type, нужно DTO
	DeletePost(ctx context.Context, ID string) error
	UpVotePost(ctx context.Context, id string) (*entities.Post, error)
	DownVotePost(ctx context.Context, id string) (*entities.Post, error)
	UnVotePost(ctx context.Context, id string) (*entities.Post, error)
}

type PostService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) *PostService {
	return &PostService{
		repo: repo,
	}
}

func (p PostService) Posts(ctx context.Context) ([]*entities.Post, error) {
	return p.repo.GetAllPosts(ctx)
}

func (p PostService) PostsByCategory(ctx context.Context, category string) ([]*entities.Post, error) {
	categoryList, err := p.repo.GetCategories(ctx)
	if err != nil {
		logger.Errorf("service.post.postbycategory.categorylist", err.Error())

		return nil, fmt.Errorf("failed to get categories: %w", err)
	}

	if !categoryExists(category, categoryList) {
		logger.Info("service.post.postbycategory.PostsByCategory: category doesn't exists")

		return nil, fmt.Errorf("category doesn't exists")
	}

	posts, err := p.repo.GetPostsByCategory(ctx, category)
	if err != nil {
		logger.Errorf("service.post.postbycategory.GetPostsByCategory", err.Error())

		return nil, fmt.Errorf("failed to get post by categories: %w", err)
	}

	return posts, nil

}

func (p PostService) PostsByUser(ctx context.Context, user entities.User) ([]*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostService) PostById(ctx context.Context, postID string) (*entities.Post, error) {
	return p.repo.PostById(ctx, postID)
}

func (p PostService) CreatePost(ctx context.Context, item entities.Post, author entities.User) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostService) DeletePost(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}

func (p PostService) UpVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostService) DownVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p PostService) UnVote(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func categoryExists(category string, categoryList []string) bool {
	exists := false
	for _, v := range categoryList {
		if v == category {
			exists = true
			break
		}
	}
	return exists
}
