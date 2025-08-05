package service

import (
	"context"
	"topic-service/internal/topic/dto/response"
	"topic-service/internal/topic/mapper"
	"topic-service/internal/topic/model"
	"topic-service/internal/topic/repository"
)

type TopicService interface {
	CreateTopic(ctx context.Context, topic *model.Topic) (*response.TopicResponse, error)
	GetTopicByID(ctx context.Context, id string) (*response.TopicResponse, error)
	UpdateTopic(ctx context.Context, id string, topic *model.Topic) error
	DeleteTopic(ctx context.Context, id string) error
	ListTopics(ctx context.Context) ([]response.TopicResponse, error)
}

type topicService struct {
	repo repository.TopicRepository
}

func NewTopicService(repo repository.TopicRepository) TopicService {
	return &topicService{repo: repo}
}

func (s *topicService) CreateTopic(ctx context.Context, topic *model.Topic) (*response.TopicResponse, error) {
	createdTopic, err := s.repo.Create(ctx, topic)
	if err != nil {
		return nil, err
	}
	return mapper.MapTopicToResponse(createdTopic), nil
}

func (s *topicService) GetTopicByID(ctx context.Context, id string) (*response.TopicResponse, error) {
	topic, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.MapTopicToResponse(topic), nil
}

func (s *topicService) UpdateTopic(ctx context.Context, id string, topic *model.Topic) error {
	return s.repo.Update(ctx, id, topic)
}

func (s *topicService) DeleteTopic(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *topicService) ListTopics(ctx context.Context) ([]response.TopicResponse, error) {
	topics, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.MapTopicsToResponses(topics), nil
}
