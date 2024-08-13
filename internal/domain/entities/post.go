package entities

// pure business logic
type Post struct {
	Category         string           `json:"category"`
	Text             string           `json:"text"`
	Title            string           `json:"title"`
	Type             string           `json:"type"`
	URL              string           `json:"url"`
	Views            uint32           `json:"views"`
	Created          string           `json:"created"`
	Author           Author           `json:"author"`
	Score            int              `json:"score"`
	UpvotePercentage int              `json:"upvotePercentage"`
	Votes            []*Vote          `json:"votes"`
	Comments         []*CommentExtend `json:"comments"`
	CommentLastID    uint32           `json:"-"`
}

// business logic with implementation logic
type PostExtend struct {
	Post `bson:"post"`
	ID   string `json:"id" bson:"id"`
}

func NewPostExtend(post Post, id string) PostExtend {
	return PostExtend{
		Post: post,
		ID:   id,
	}
}

type Author struct {
	Username string `json:"username"`
	ID       string `json:"id"`
}

func NewAuthor() *Author {
	return &Author{}
}

type Vote struct {
	UserID string `json:"user"`
	Vote   int    `json:"vote"`
}

func NewVote() *Vote {
	return &Vote{}
}
