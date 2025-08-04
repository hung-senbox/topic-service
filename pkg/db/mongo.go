package db

import (
	"context"
	"fmt"
	"log"
	"term-service/pkg/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var TermCollection *mongo.Collection

func ConnectMongoDB() {
	d := config.AppConfig.Database.Mongo

	var uri string
	if d.User != "" && d.Password != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s", d.User, d.Password, d.Host, d.Port)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s", d.Host, d.Port)
	}

	clientOptions := options.Client().ApplyURI(uri)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	MongoClient, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	if err := MongoClient.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB ping failed: %v", err)
	}

	TermCollection = MongoClient.Database(d.Name).Collection("terms")
	log.Println("Connected to MongoDB and loaded 'terms' collection")
}
