package memory

import (
	s "redditClone/internal/storage"
)

type ItemMemoryRepository struct {
	lastID uint32
	data   []*s.Item
}

var _ s.ItemsRepo = &ItemMemoryRepository{}

func NewMemoryRepo() *ItemMemoryRepository {
	res := &ItemMemoryRepository{
		data: make([]*s.Item, 0, 10),
	}
	itemTest := []*s.Item{
		{
			Category: "music",
			Text:     "text",
			Title:    "TEST",
			Type:     "text",

			ID:      "656b54d31d06de00132f7ddc",
			URL:     "-",
			Views:   324,
			Created: "023-12-02T16:01:23.248Z",
			Author:  s.Author{Username: "CHAPA", ID: "228"},

			Score:            22,
			UpvotePercentage: 78,
			Votes:            []*s.Vote{{User: "bibp", Vote: 1}, {User: "boba", Vote: -1}},

			Comments: []*s.Comment{{Author: s.Author{Username: "CHAPA", ID: "228"}, Body: "lksdjf", Created: "dsklfjsldkfjsldkf", ID: "1"}},
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
			Author:  s.Author{Username: "CHAPA", ID: "228"},

			Score:            23,
			UpvotePercentage: 78,
			Votes:            []*s.Vote{{User: "bibp", Vote: 1}, {User: "boba", Vote: -1}},

			Comments: []*s.Comment{{Author: s.Author{Username: "CHAPA", ID: "228"}, Body: "lksdjf", Created: "dsklfjsldkfjsldkf", ID: "1"}},
		},
	}
	res.data = append(res.data, itemTest...)
	return res
}

func (i *ItemMemoryRepository) GetAll() ([]*s.Item, error) {
	// TODO: Add sorting
	return i.data, nil
}

func (i *ItemMemoryRepository) GetByCategory(category string) ([]*s.Item, error) {
	// TODO: Add sorting
	dataByCategory := make([]*s.Item, 0)
	for _, v := range i.data {
		if v.Category == category {
			dataByCategory = append(dataByCategory, v)
		}
	}
	return dataByCategory, nil
}

func (i *ItemMemoryRepository) GetPost(id string) (*s.Item, error) {
	// TODO: Add sorting
	var post *s.Item
	var err error
	for _, v := range i.data {
		if v.ID == id {
			post = v
		}
	}

	if post == nil {
		err = s.ErrNotFound
	}

	return post, err
}
