package inMemory

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"time"
)

// реализуем интерфейс storage

type Posts struct {
	lastID uint32
	data   []*entities.PostExtend
}

var _ interfaces.IPostRepository = &Posts{}

func NewPosts() *Posts {
	initPosts := []*entities.PostExtend{
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
	data := make([]*entities.PostExtend, 0, 10)
	return &Posts{
		lastID: 2,
		data:   append(data, initPosts...),
	}
}

func (p Posts) Add(ctx context.Context, post entities.PostExtend) error {

	panic("implement me")
}

func (p Posts) Get(ctx context.Context, postID string) (entities.PostExtend, error) {

	panic("implement me")
}

func (p Posts) GetWhereCategory(ctx context.Context, category string) ([]entities.PostExtend, error) {

	panic("implement me")
}

func (p Posts) GetWhereUsername(ctx context.Context, username string) ([]entities.PostExtend, error) {

	panic("implement me")
}

func (p Posts) GetAll(ctx context.Context) ([]entities.PostExtend, error) {

	panic("implement me")
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
