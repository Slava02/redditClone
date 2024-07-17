package storage

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrExists   = errors.New("exists")
)

type ItemsRepo interface {
	GetAll() ([]*Item, error)
	GetByCategory(category string) ([]*Item, error)
	GetPost(id string) (*Item, error)
}

type Comment struct {
	Author  Author `json:"author"`
	Body    string `json:"body"`
	Created string `json:"created"`
	ID      string `json:"id"`
}

func NewComment() *Comment {
	return &Comment{}
}

type Vote struct {
	User string `json:"user"`
	Vote int    `json:"vote"`
}

func NewVote() *Vote {
	return &Vote{}
}

type Author struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

func NewAuthor() *Author {
	return &Author{}
}

type Item struct {
	Category         string     `json:"category"`
	Text             string     `json:"text"`
	Title            string     `json:"title"`
	Type             string     `json:"type"`
	ID               string     `json:"id"`
	URL              string     `json:"url"`
	Views            uint32     `json:"views"`
	Created          string     `json:"created"`
	Author           Author     `json:"author"`
	Score            int        `json:"score"`
	UpvotePercentage int        `json:"upvotePercentage"`
	Votes            []*Vote    `json:"votes"`
	Comments         []*Comment `json:"comments"`
	CommentLastID    uint32     `json:"-"`
}

func NewItem() *Item {
	return &Item{}
}
