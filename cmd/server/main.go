package main

import (
	"log"
	"os"

	// "os"

	"topic-service/pkg/config"
	"topic-service/pkg/consul"
	"topic-service/pkg/db"
	"topic-service/pkg/router"

	"topic-service/pkg/zap"
)

func main() {
	filePath := os.Args[1]
	if filePath == "" {
		filePath = "configs/config.yaml"
	}

	config.LoadConfig(filePath)

	cfg := config.AppConfig

	// logger.WriteLogData("info", map[string]any{"id": 123, "name": "Hung"})

	//logger
	logger, err := zap.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	//consul
	consulConn := consul.NewConsulConn(logger, cfg)
	consulConn.Connect()
	defer consulConn.Deregister()

	//db
	db.ConnectMongoDB()

	r := router.SetupRouter(db.TopicCollection)
	port := cfg.Server.Port
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to run server:", err)
	}
}
