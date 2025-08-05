package handler

import (
	"net/http"

	"topic-service/internal/topic/dto/response"
	"topic-service/internal/topic/model"
	"topic-service/internal/topic/service"

	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	service service.TopicService
}

func NewTopicHandler(service service.TopicService) *TopicHandler {
	return &TopicHandler{service: service}
}

// POST /topics
func (h *TopicHandler) CreateTopic(c *gin.Context) {
	var topic model.Topic
	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	result, err := h.service.CreateTopic(c.Request.Context(), &topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to create topic",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, response.SucceedResponse{
		Code:    http.StatusCreated,
		Message: "Topic created successfully",
		Data:    result,
	})
}

// GET /topics/:id
func (h *TopicHandler) GetTopicByID(c *gin.Context) {
	id := c.Param("id")

	topic, err := h.service.GetTopicByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.FailedResponse{
			Code:    http.StatusNotFound,
			Message: "Topic not found",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "Topic retrieved successfully",
		Data:    topic,
	})
}

// PUT /topics/:id
func (h *TopicHandler) UpdateTopic(c *gin.Context) {
	id := c.Param("id")
	var topic model.Topic

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, response.FailedResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid request body",
			Error:   err.Error(),
		})
		return
	}

	err := h.service.UpdateTopic(c.Request.Context(), id, &topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to update topic",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "Topic updated successfully",
		Data:    nil,
	})
}

// DELETE /topics/:id
func (h *TopicHandler) DeleteTopic(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteTopic(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.FailedResponse{
			Code:    http.StatusNotFound,
			Message: "Topic not found",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "Topic deleted successfully",
		Data:    nil,
	})
}

// GET /topics
func (h *TopicHandler) ListTopics(c *gin.Context) {
	topics, err := h.service.ListTopics(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.FailedResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to list topics",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response.SucceedResponse{
		Code:    http.StatusOK,
		Message: "Topics retrieved successfully",
		Data:    topics,
	})
}
