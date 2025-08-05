package repository

import (
	"context"
	"errors"
	"time"
	"topic-service/internal/topic/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopicRepository interface {
	Create(ctx context.Context, topic *model.Topic) (*model.Topic, error)
	GetByID(ctx context.Context, id string) (*model.Topic, error)
	Update(ctx context.Context, id string, topic *model.Topic) error
	Delete(ctx context.Context, id string) error
	GetAll(ctx context.Context) ([]*model.Topic, error)
}

type topicRepository struct {
	collection *mongo.Collection
}

func NewTopicRepository(collection *mongo.Collection) TopicRepository {
	return &topicRepository{collection}
}

func (r *topicRepository) Create(ctx context.Context, topic *model.Topic) (*model.Topic, error) {

	_, err := r.collection.InsertOne(ctx, topic)
	if err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *topicRepository) GetByID(ctx context.Context, id string) (*model.Topic, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var topic model.Topic
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&topic)
	if err != nil {
		return nil, err
	}
	return &topic, nil
}

func (r *topicRepository) Update(ctx context.Context, id string, updated *model.Topic) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	updated.UpdatedAt = time.Now()

	update := bson.M{
		"$set": bson.M{
			"title":      updated.Title,
			"updated_at": updated.UpdatedAt,
		},
	}

	result, err := r.collection.UpdateByID(ctx, objectID, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *topicRepository) Delete(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}

func (r *topicRepository) GetAll(ctx context.Context) ([]*model.Topic, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var topics []*model.Topic
	for cursor.Next(ctx) {
		var topic model.Topic
		if err := cursor.Decode(&topic); err != nil {
			return nil, err
		}
		topics = append(topics, &topic)
	}
	return topics, nil
}
