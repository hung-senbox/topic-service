package handler

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateTopic(c.Request.Context(), &topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create topic"})
		return
	}

	c.JSON(http.StatusCreated, result)
}

// GET /topics/:id
func (h *TopicHandler) GetTopicByID(c *gin.Context) {
	id := c.Param("id")

	topic, err := h.service.GetTopicByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "topic not found"})
		return
	}

	c.JSON(http.StatusOK, topic)
}

// PUT /topics/:id
func (h *TopicHandler) UpdateTopic(c *gin.Context) {
	id := c.Param("id")
	var topic model.Topic

	if err := c.ShouldBindJSON(&topic); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.service.UpdateTopic(c.Request.Context(), id, &topic)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update topic"})
		return
	}

	c.Status(http.StatusNoContent)
}

// DELETE /topics/:id
func (h *TopicHandler) DeleteTopic(c *gin.Context) {
	id := c.Param("id")

	err := h.service.DeleteTopic(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "topic not found"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GET /topics
func (h *TopicHandler) ListTopics(c *gin.Context) {
	topics, err := h.service.ListTopics(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list topics"})
		return
	}

	c.JSON(http.StatusOK, topics)
}
