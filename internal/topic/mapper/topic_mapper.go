package mapper

import (
	"topic-service/internal/topic/dto/response"
	"topic-service/internal/topic/model"
)

// Mapper: Topic model -> TopicResponse
func MapTopicToResponse(t *model.Topic) *response.TopicResponse {
	if t == nil {
		return nil
	}

	return &response.TopicResponse{
		ID:        t.ID.Hex(),
		Title:     t.Title,
		Icon:      t.Icon,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}
}

func MapTopicsToResponses(topics []*model.Topic) []response.TopicResponse {
	var responses []response.TopicResponse
	for _, t := range topics {
		res := MapTopicToResponse(t)
		if res != nil {
			responses = append(responses, *res)
		}
	}
	return responses
}
