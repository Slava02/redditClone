package inMemory

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
	"sync"
	"time"
)

// реализуем интерфейс storage

type Posts struct {
	lastID uint32
	mutex  sync.RWMutex
	data   []entities.PostExtend
}

var _ interfaces.IPostRepository = &Posts{}

func NewPosts() *Posts {
	initPosts := []entities.PostExtend{
		{
			ID: "656b54d31d06de00132f7ddc",
			Post: entities.Post{
				Category: "music",
				Text:     "text",
				Title:    "TEST",
				Type:     "text",

				URL:     "-",
				Views:   324,
				Created: "023-12-02T16:01:23.248Z",
				Author:  entities.Author{Username: "CHAPA", ID: "228"},

				Score:            22,
				UpvotePercentage: 78,
				Votes:            []*entities.Vote{{UserID: "bibp", Vote: 1}, {UserID: "boba", Vote: -1}},

				Comments: []*entities.CommentExtend{
					{
						ID: "1",
						Comment: entities.Comment{
							Author: entities.Author{
								Username: "CHAPA",
								ID:       "228",
							},
							Body:    "lksdjf",
							Created: time.Now(),
						},
					},
				},
			},
		},
	}
	data := make([]entities.PostExtend, 0, 10)
	return &Posts{
		lastID: 2,
		data:   append(data, initPosts...),
	}
}

func (p Posts) Add(ctx context.Context, post entities.PostExtend) error {

	panic("implement me")
}

func (p Posts) Get(ctx context.Context, postID string) (entities.PostExtend, error) {
	const op = "repo.post.Get: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for _, v := range p.data {
		if v.ID == postID {
			return v, nil
		}
	}

	logger.Info(op, fmt.Sprintf("POST NOT FOUND (postID: %s)", postID))

	return entities.PostExtend{}, fmt.Errorf("%w", repository.ErrNotFound)
}

func (p Posts) GetWhereCategory(ctx context.Context, category string) ([]entities.PostExtend, error) {
	const op = "repo.post.GetWhereCategory: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	postByCategory := make([]entities.PostExtend, 0)
	for _, v := range p.data {
		if v.Category == category {
			postByCategory = append(postByCategory, v)
		}
	}

	return postByCategory, nil
}

func (p Posts) GetWhereUsername(ctx context.Context, username string) ([]entities.PostExtend, error) {

	panic("implement me")
}

func (p Posts) GetAll(ctx context.Context) ([]entities.PostExtend, error) {
	const op = "repo.post.GetAll: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.data, nil
}

func (p Posts) Update(ctx context.Context, postID string, newPost entities.PostExtend) error {

	panic("implement me")
}

func (p Posts) Delete(ctx context.Context, postID string) error {

	panic("implement me")
}

func (p Posts) AddComment(ctx context.Context, postID string, comment entities.CommentExtend) (entities.PostExtend, error) {

	panic("implement me")
}

func (p Posts) GetComment(ctx context.Context, postID string, commentID string) (entities.CommentExtend, error) {

	panic("implement me")
}

func (p Posts) DeleteComment(ctx context.Context, postID string, commentID string) (entities.PostExtend, error) {

	panic("implement me")
}
