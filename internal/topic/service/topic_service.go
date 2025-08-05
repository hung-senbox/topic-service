package service

import (
	"context"
	"topic-service/internal/topic/model"
	"topic-service/internal/topic/repository"
)

type TopicService interface {
	CreateTopic(ctx context.Context, topic *model.Topic) (*model.Topic, error)
	GetTopicByID(ctx context.Context, id string) (*model.Topic, error)
	UpdateTopic(ctx context.Context, id string, topic *model.Topic) error
	DeleteTopic(ctx context.Context, id string) error
	ListTopics(ctx context.Context) ([]*model.Topic, error)
}

type topicService struct {
	repo repository.TopicRepository
}

func NewTopicService(repo repository.TopicRepository) TopicService {
	return &topicService{repo: repo}
}

func (s *topicService) CreateTopic(ctx context.Context, topic *model.Topic) (*model.Topic, error) {
	return s.repo.Create(ctx, topic)
}

func (s *topicService) GetTopicByID(ctx context.Context, id string) (*model.Topic, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *topicService) UpdateTopic(ctx context.Context, id string, topic *model.Topic) error {
	return s.repo.Update(ctx, id, topic)
}

func (s *topicService) DeleteTopic(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *topicService) ListTopics(ctx context.Context) ([]*model.Topic, error) {
	return s.repo.GetAll(ctx)
}
