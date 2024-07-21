package entities

type Post struct {
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

func NewPost() *Post {
	return &Post{}
}

type Author struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

func NewAuthor() *Author {
	return &Author{}
}

type Vote struct {
	User string `json:"user"`
	Vote int    `json:"vote"`
}

func NewVote() *Vote {
	return &Vote{}
}
