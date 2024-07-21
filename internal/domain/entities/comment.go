package entities

type Comment struct {
	Author  Author `json:"author"`
	Body    string `json:"body"`
	Created string `json:"created"`
	ID      string `json:"id"`
}

func NewComment() *Comment {
	return &Comment{}
}
