package entities

import (
	"errors"
	"fmt"
)

var (
	ErrAlreadyUpvote   = errors.New("post has been already upvoted")
	ErrAlreadyDownvote = errors.New("post has been already downvoted")
	ErrAlreadyUnvote   = errors.New("post has been already unvoted")
)

// pure business logic
type Post struct {
	Category         string          `json:"category"`
	Text             string          `json:"text"`
	Title            string          `json:"title"`
	Type             string          `json:"type"`
	URL              string          `json:"url"`
	Views            uint32          `json:"views"`
	Created          string          `json:"created"`
	Author           Author          `json:"author"`
	Score            int             `json:"score"`
	UpvotePercentage int             `json:"upvotePercentage"`
	Votes            []Vote          `json:"votes"`
	Comments         []CommentExtend `json:"comments"`
	CommentLastID    uint32          `json:"-"`
}

func NewPost(category, text, title, postType, url, created string, author Author) Post {
	return Post{
		Category: category,
		Title:    title,
		Text:     text,
		Type:     postType,
		URL:      url,
		Author:   author,
		Votes:    make([]Vote, 0),
		Comments: make([]CommentExtend, 0),
		Views:    1,
		Created:  created,
	}
}

func (post *Post) View() {
	post.Views++
	return
}

func (post *Post) Upvote(userID string) error {
	for i := 0; i < len(post.Votes); i++ {
		if post.Votes[i].UserID == userID {
			if post.Votes[i].Vote == 1 {
				return fmt.Errorf("[Post.Upvote]: %w", ErrAlreadyUpvote)
			}
			copy(post.Votes[i:], post.Votes[i+1:])
			post.Votes[len(post.Votes)-1] = Vote{
				UserID: userID,
				Vote:   +1,
			}
			post.Score += 2
			post.updateScore()
			return nil
		}
	}

	post.Votes = append(post.Votes, Vote{
		UserID: userID,
		Vote:   1,
	})
	post.Score++
	post.updateScore()

	return nil

}

func (post *Post) Downvote(userID string) error {
	for i := 0; i < len(post.Votes); i++ {
		if post.Votes[i].UserID == userID {
			if post.Votes[i].Vote == -1 {
				return fmt.Errorf("[Post.Downvote]: %w", ErrAlreadyDownvote)
			}
			copy(post.Votes[i:], post.Votes[i+1:])
			post.Votes[len(post.Votes)-1] = Vote{
				UserID: userID,
				Vote:   -1,
			}
			post.Score -= 2
			post.updateScore()

			return nil
		}
	}

	post.Votes = append(post.Votes, Vote{
		UserID: userID,
		Vote:   -1,
	})
	post.Score--
	post.updateScore()

	return nil
}

func (post *Post) Unvote(userID string) error {
	for i := 0; i < len(post.Votes); i++ {
		if post.Votes[i].UserID == userID {
			if post.Votes[i].Vote == 1 {
				post.Score--
			} else {
				post.Score++
			}
			copy(post.Votes[i:], post.Votes[i+1:])
			post.Votes[len(post.Votes)-1] = Vote{}
			post.Votes = post.Votes[:len(post.Votes)-1]
			post.updateScore()

			return nil
		}
	}

	return fmt.Errorf("[Post.Unvote]: %w", ErrAlreadyUnvote)
}

func (post *Post) updateScore() {
	if post.Score > 0 {
		post.UpvotePercentage = int(float64((post.Score+len(post.Votes))/2) * 100 / float64(len(post.Votes)))
	} else {
		post.UpvotePercentage = 0
	}
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

func NewAuthor(id, username string) Author {
	return Author{
		Username: username,
		ID:       id,
	}
}

type Vote struct {
	UserID string `json:"user"`
	Vote   int    `json:"vote"`
}

func NewVote() *Vote {
	return &Vote{}
}
