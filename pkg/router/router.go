package router

import (
	"topic-service/internal/topic/handler"
	"topic-service/internal/topic/repository"
	"topic-service/internal/topic/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(topicCol *mongo.Collection) *gin.Engine {
	r := gin.Default()

	topicRepo := repository.NewTopicRepository(topicCol)
	topicSvc := service.NewTopicService(topicRepo)
	topicHandler := handler.NewTopicHandler(topicSvc)

	v1 := r.Group("/api/v1")
	{
		topicGroup := v1.Group("/topics")
		{
			topicGroup.POST("", topicHandler.CreateTopic)
			topicGroup.GET("/:id", topicHandler.GetTopicByID)
			topicGroup.PUT("/:id", topicHandler.UpdateTopic)
			topicGroup.DELETE("/:id", topicHandler.DeleteTopic)
			topicGroup.GET("", topicHandler.ListTopics)
		}
	}

	return r
}
