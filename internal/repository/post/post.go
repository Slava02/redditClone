package post

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"redditClone/internal/domain/entities"
	"redditClone/internal/interfaces"
	"redditClone/internal/repository"
)

type PostRepoMongoDb struct {
	client     *mongo.Client
	collection *mongo.Collection
}

var _ interfaces.IPostRepository = (*PostRepoMongoDb)(nil)

func NewPosts(client *mongo.Client, col *mongo.Collection) *PostRepoMongoDb {
	return &PostRepoMongoDb{
		client:     client,
		collection: col,
	}
}

func (p PostRepoMongoDb) Add(ctx context.Context, post entities.PostExtend) error {
	const op = "repository.mongodb.posts.Add: "

	_, err := p.collection.InsertOne(ctx, post)
	if err != nil {
		return fmt.Errorf("%s:%w", op, err)
	}

	return nil
}

func (p PostRepoMongoDb) Get(ctx context.Context, postID string) (entities.PostExtend, error) {
	const op = "repository.mongodb.posts.Get: "

	filter := bson.D{{"id", postID}}

	res := p.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return entities.PostExtend{}, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		} else {
			return entities.PostExtend{}, fmt.Errorf("%s: %w", op, res.Err())
		}
	}

	var post entities.PostExtend
	err := res.Decode(&post)
	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%s: %w", op, err)
	}

	return post, nil
}

func (p PostRepoMongoDb) GetWhereCategory(ctx context.Context, category string) ([]entities.PostExtend, error) {
	const op = "repository.mongodb.posts.GetWhereCategory: "

	filter := bson.D{{"category", category}}

	c, err := p.collection.Find(ctx, filter)
	if c.Err() != nil {
		if errors.Is(c.Err(), mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		} else {
			return nil, fmt.Errorf("%s: %w", op, c.Err())
		}
	}

	var posts []entities.PostExtend
	err = c.All(ctx, &posts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

func (p PostRepoMongoDb) GetWhereUsername(ctx context.Context, username string) ([]entities.PostExtend, error) {
	const op = "repository.mongodb.posts.GetWhereUsername: "

	filter := bson.D{{"username", username}}

	c, err := p.collection.Find(ctx, filter)
	if c.Err() != nil {
		if errors.Is(c.Err(), mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		} else {
			return nil, fmt.Errorf("%s: %w", op, c.Err())
		}
	}

	var posts []entities.PostExtend
	err = c.All(ctx, &posts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

func (p PostRepoMongoDb) GetAll(ctx context.Context) ([]entities.PostExtend, error) {
	const op = "repository.mongodb.posts.GetAll: "

	c, err := p.collection.Find(ctx, bson.M{})
	if c.Err() != nil {
		if errors.Is(c.Err(), mongo.ErrNoDocuments) {
			return nil, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		} else {
			return nil, fmt.Errorf("%s: %w", op, c.Err())
		}
	}

	var posts []entities.PostExtend
	err = c.All(ctx, &posts)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

func (p PostRepoMongoDb) Update(ctx context.Context, postID string, newPost entities.PostExtend) error {
	const op = "repository.mongodb.posts.Update: "

	filter := bson.D{{"id", postID}}

	res, err := p.collection.UpdateOne(ctx, filter, newPost)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.MatchedCount == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}

func (p PostRepoMongoDb) Delete(ctx context.Context, postID string) error {
	const op = "repository.mongodb.posts.Delete: "

	filter := bson.D{{"id", postID}}

	res, err := p.collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if res.DeletedCount == 0 {
		return fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	return nil
}

func (p PostRepoMongoDb) AddComment(ctx context.Context, postID string, comment entities.CommentExtend) (entities.PostExtend, error) {
	const op = "repository.mongodb.posts.AddComment: "

	filter := bson.D{{"id", postID}}
	add := bson.M{"$push": bson.M{"post.$.comments": comment}}

	res, err := p.collection.UpdateOne(ctx, filter, add)
	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%s: %w", op, err)
	} else if res.MatchedCount == 0 {
		return entities.PostExtend{}, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
	}

	post, err := p.Get(ctx, postID)
	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%s:%w", op, err)
	}

	return post, nil
}

func (p PostRepoMongoDb) GetComment(ctx context.Context, postID string, commentID string) (entities.CommentExtend, error) {
	const op = "repository.mongodb.posts.GetComment: "

	filter := bson.D{{"id", postID}}

	res := p.collection.FindOne(ctx, filter)
	if res.Err() != nil {
		if errors.Is(res.Err(), mongo.ErrNoDocuments) {
			return entities.PostExtend{}, fmt.Errorf("%s: %w", op, repository.ErrNotFound)
		} else {
			return entities.PostExtend{}, fmt.Errorf("%s: %w", op, res.Err())
		}
	}

	var post entities.PostExtend
	err := res.Decode(&post)
	if err != nil {
		return entities.PostExtend{}, fmt.Errorf("%s: %w", op, err)
	}

	return post, nil
}

func (p PostRepoMongoDb) DeleteComment(ctx context.Context, postID string, commentID string) (entities.PostExtend, error) {
	//TODO implement me
	panic("implement me")
}
