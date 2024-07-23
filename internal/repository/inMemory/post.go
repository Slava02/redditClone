package inMemory

import (
	"context"
	"redditClone/internal/domain/entities"
	"redditClone/internal/domain/service"
	"redditClone/internal/repository"
)

// реализуем интерфейс storage

type Posts struct {
	lastID uint32
	data   []*entities.Post
}

var _ service.PostRepository = &Posts{}

func NewPosts() *Posts {
	initPosts := []*entities.Post{
		{
			Category: "music",
			Text:     "text",
			Title:    "TEST",
			Type:     "text",

			ID:      "656b54d31d06de00132f7ddc",
			URL:     "-",
			Views:   324,
			Created: "023-12-02T16:01:23.248Z",
			Author:  entities.Author{Username: "CHAPA", ID: "228"},

			Score:            22,
			UpvotePercentage: 78,
			Votes:            []*entities.Vote{{User: "bibp", Vote: 1}, {User: "boba", Vote: -1}},

			Comments: []*entities.Comment{{Author: entities.Author{Username: "CHAPA", ID: "228"}, Body: "lksdjf", Created: "dsklfjsldkfjsldkf", ID: "1"}},
		},
		{
			Category: "programming",
			Text:     "text",
			Title:    "TEST LINK",
			Type:     "link",

			ID:      "656b54d31d06de00134f7ddc",
			URL:     "ya.com",
			Views:   324,
			Created: "023-12-02T17:01:23.248Z",
			Author:  entities.Author{Username: "CHAPA", ID: "228"},

			Score:            23,
			UpvotePercentage: 78,
			Votes:            []*entities.Vote{{User: "bibp", Vote: 1}, {User: "boba", Vote: -1}},

			Comments: []*entities.Comment{{Author: entities.Author{Username: "CHAPA", ID: "228"}, Body: "lksdjf", Created: "dsklfjsldkfjsldkf", ID: "1"}},
		},
	}
	data := make([]*entities.Post, 0, 10)
	return &Posts{
		lastID: 2,
		data:   append(data, initPosts...),
	}
}

func (p *Posts) GetAllPosts(ctx context.Context) ([]*entities.Post, error) {
	return p.data, nil
}

func (p *Posts) GetPostsByCategory(ctx context.Context, category string) ([]*entities.Post, error) {
	//TODO add sorting
	dataByCategory := make([]*entities.Post, 0)
	for _, v := range p.data {
		if v.Category == category {
			dataByCategory = append(dataByCategory, v)
		}
	}
	return dataByCategory, nil
}

func (p *Posts) GetCategories(ctx context.Context) ([]string, error) {
	return CategoriesList, nil
}

func (p *Posts) PostsByUser(ctx context.Context, user entities.User) ([]*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Posts) PostById(ctx context.Context, postID string) (*entities.Post, error) {
	for _, v := range p.data {
		if v.ID == postID {
			return v, nil
		}
	}

	return nil, repository.ErrNotFound
}

func (p *Posts) CreatePost(ctx context.Context, item entities.Post, author entities.Author) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Posts) DeletePost(ctx context.Context, ID string) error {
	//TODO implement me
	panic("implement me")
}

func (p *Posts) UpVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Posts) DownVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Posts) UnVotePost(ctx context.Context, id string) (*entities.Post, error) {
	//TODO implement me
	panic("implement me")
}

var CategoriesList = []string{
	"music",
	"funny",
	"videos",
	"programming",
	"news",
	"fashion",
}
