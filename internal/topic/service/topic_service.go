package service

import (
	"context"
	"errors"
	"fmt"
	"topic-service/internal/gateway"
	"topic-service/internal/topic/dto/response"
	"topic-service/internal/topic/mapper"
	"topic-service/internal/topic/model"
	"topic-service/internal/topic/repository"

	"github.com/gin-gonic/gin"
)

type TopicService interface {
	CreateTopic(ctx *gin.Context, topic *model.Topic) (*response.TopicResponse, error)
	GetTopicByID(ctx context.Context, id string) (*response.TopicResponse, error)
	UpdateTopic(ctx context.Context, id string, topic *model.Topic) error
	DeleteTopic(ctx context.Context, id string) error
	ListTopics(ctx context.Context) ([]response.TopicResponse, error)
}

type topicService struct {
	repo        repository.TopicRepository
	userGateway gateway.UserGateway
}

func NewTopicService(repo repository.TopicRepository, userGateway gateway.UserGateway) TopicService {
	return &topicService{
		repo:        repo,
		userGateway: userGateway,
	}
}

func (s *topicService) CreateTopic(ctx *gin.Context, topic *model.Topic) (*response.TopicResponse, error) {

	userIDRaw, exists := ctx.Get("user_id")
	if !exists {
		return nil, errors.New("user ID not found in context")
	}

	userID, ok := userIDRaw.(string)
	if !ok || userID == "" {
		return nil, errors.New("user ID is not a valid string")
	}

	// Gọi gateway để kiểm tra user
	user, err := s.userGateway.GetAuthorInfo(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get author info: %w", err)
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

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

func (s *topicService) GetAuthorInfo(ctx context.Context, userID string) (*gateway.User, error) {
	return s.userGateway.GetAuthorInfo(ctx, userID)
}
