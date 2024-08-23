package entities

import "time"

// pure business logic
type Comment struct {
	Author  Author    `json:"author" bson:"author"`
	Body    string    `json:"body" bson:"body"`
	Created time.Time `json:"created" bson:"created"`
}

// business logic with implementation logic
type CommentExtend struct {
	Comment
	ID string `json:"id" bson:"id"`
}

func NewCommentExtend(comment Comment, id string) CommentExtend {
	return CommentExtend{
		Comment: comment,
		ID:      id,
	}
}
