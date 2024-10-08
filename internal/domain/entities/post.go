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
	Category         string          `json:"category" bson:"category"`
	Text             string          `json:"text" bson:"text"`
	Title            string          `json:"title" bson:"title"`
	Type             string          `json:"type" bson:"type"`
	URL              string          `json:"url" bson:"URL"`
	Views            uint32          `json:"views" bson:"views"`
	Created          string          `json:"created" bson:"created"`
	Author           Author          `json:"author" bson:"author"`
	Score            int             `json:"score" bson:"score"`
	UpvotePercentage int             `json:"upvotePercentage" bson:"upvotePercentage"`
	Votes            []Vote          `json:"votes" bson:"votes"`
	Comments         []CommentExtend `json:"comments" bson:"comments"`
	CommentLastID    uint32          `json:"-" bson:"-"`
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
	Username string `json:"username" bson:"username"`
	ID       string `json:"id" bson:"id"`
}

func NewAuthor(id, username string) Author {
	return Author{
		Username: username,
		ID:       id,
	}
}

type Vote struct {
	UserID string `json:"user" bson:"userID"`
	Vote   int    `json:"vote" bson:"vote"`
}

func NewVote() *Vote {
	return &Vote{}
}
