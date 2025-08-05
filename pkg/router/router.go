package router

import (
	"topic-service/internal/gateway"
	"topic-service/internal/topic/handler"
	"topic-service/internal/topic/middleware"
	"topic-service/internal/topic/repository"
	"topic-service/internal/topic/service"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(topicCol *mongo.Collection) *gin.Engine {
	r := gin.Default()

	// Init Consul client
	consulClient, _ := api.NewClient(api.DefaultConfig())

	// Tạo UserGateway
	userGateway := gateway.NewUserGateway("go-main-service", consulClient)

	// Init repository và service
	topicRepo := repository.NewTopicRepository(topicCol)
	topicSvc := service.NewTopicService(topicRepo, userGateway)
	topicHandler := handler.NewTopicHandler(topicSvc)

	v1 := r.Group("/api/v1")
	{
		topicGroup := v1.Group("/topic", middleware.Secured())
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
