package inMemory

import (
	"context"
	"fmt"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/internal/repository"
	"redditClone/pkg/logger"
	"sync"
)

// реализуем интерфейс storage

type Posts struct {
	lastID uint32
	mutex  sync.RWMutex
	data   []entities.PostExtend
}

var _ interfaces.IPostRepository = &Posts{}

func NewPosts() *Posts {
	//initPosts := []entities.PostExtend{
	//	{
	//		ID: "656b54d31d06de00132f7ddc",
	//		Post: entities.Post{
	//			Category: "music",
	//			Text:     "text",
	//			Title:    "TEST",
	//			Type:     "text",
	//
	//			URL:     "-",
	//			Views:   324,
	//			Created: "023-12-02T16:01:23.248Z",
	//			Author:  entities.Author{Username: "CHAPA", ID: "228"},
	//
	//			Score:            22,
	//			UpvotePercentage: 78,
	//			Votes:            []*entities.Vote{{UserID: "bibp", Vote: 1}, {UserID: "boba", Vote: -1}},
	//
	//			Comments: []*entities.CommentExtend{
	//				{
	//					ID: "1",
	//					Comment: entities.Comment{
	//						Author: entities.Author{
	//							Username: "CHAPA",
	//							ID:       "228",
	//						},
	//						Body:    "lksdjf",
	//						Created: time.Now(),
	//					},
	//				},
	//			},
	//		},
	//	},
	//}
	data := make([]entities.PostExtend, 0)
	return &Posts{
		lastID: 1,
		data:   data,
	}
}

func (p *Posts) Add(ctx context.Context, post entities.PostExtend) error {
	const op = "repo.post.Add: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	p.data = append(p.data, post)
	p.lastID++

	return nil
}

func (p *Posts) Get(ctx context.Context, postID string) (entities.PostExtend, error) {
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

func (p *Posts) GetWhereCategory(ctx context.Context, category string) ([]entities.PostExtend, error) {
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

func (p *Posts) GetWhereUsername(ctx context.Context, username string) ([]entities.PostExtend, error) {
	const op = "repo.post.GetWhereUsername: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	postByUser := make([]entities.PostExtend, 0)
	for _, v := range p.data {
		if v.Author.Username == username {
			postByUser = append(postByUser, v)
		}
	}

	return postByUser, nil
}

func (p *Posts) GetAll(ctx context.Context) ([]entities.PostExtend, error) {
	const op = "repo.post.GetAll: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.data, nil
}

func (p *Posts) Update(ctx context.Context, postID string, newPost entities.PostExtend) error {

	panic("implement me")
}

func (p *Posts) Delete(ctx context.Context, postID string) error {
	const op = "repo.post.Get: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for i, v := range p.data {
		if v.ID == postID {
			p.data = append(p.data[:i], p.data[i+1:]...)

			return nil
		}
	}

	logger.Info(op + fmt.Sprintf("POST NOT FOUND (postID: %s)", postID))

	return fmt.Errorf("%w", repository.ErrNotFound)
}

func (p *Posts) AddComment(ctx context.Context, postID string, comment entities.CommentExtend) (entities.PostExtend, error) {
	const op = "repo.post.AddComment: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for _, v := range p.data {
		if v.ID == postID {
			v.Comments = append(v.Comments, &comment)

			return v, nil
		}
	}

	logger.Info(op + fmt.Sprintf("POST NOT FOUND (postID: %s)", postID))

	return entities.PostExtend{}, fmt.Errorf("%w", repository.ErrNotFound)
}

func (p *Posts) GetComment(ctx context.Context, postID string, commentID string) (entities.CommentExtend, error) {
	const op = "repo.post.GetComment: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for _, v := range p.data {
		if v.ID == postID {
			for _, c := range v.Comments {
				if c.ID == commentID {
					return *c, nil
				}
			}
		}
	}

	logger.Info(op, fmt.Sprintf("POST NOT FOUND (postID: %s)", postID))

	return entities.CommentExtend{}, fmt.Errorf("%w", repository.ErrNotFound)
}

func (p *Posts) DeleteComment(ctx context.Context, postID string, commentID string) (entities.PostExtend, error) {
	const op = "repo.post.DeleteComment: "

	p.mutex.RLock()
	defer p.mutex.RUnlock()

	for _, v := range p.data {
		if v.ID == postID {
			for i, c := range v.Comments {
				if c.ID == commentID {
					v.Comments = append(v.Comments[:i], v.Comments[:i+1]...)

					return v, nil
				}
			}
		}
	}

	logger.Info(op + fmt.Sprintf("POST NOT FOUND (postID: %s)", postID))

	return entities.PostExtend{}, fmt.Errorf("%w", repository.ErrNotFound)
}
