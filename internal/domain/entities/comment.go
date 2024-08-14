package entities

import "time"

// pure business logic
type Comment struct {
	Author  Author    `json:"author"`
	Body    string    `json:"body"`
	Created time.Time `json:"created"`
}

// business logic with implementation logic
type CommentExtend struct {
	Comment
	ID string `json:"id"`
}

func NewCommentExtend(comment Comment, id string) CommentExtend {
	return CommentExtend{
		Comment: comment,
		ID:      id,
	}
}
