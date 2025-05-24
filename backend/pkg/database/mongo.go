package database

import (
	"context"
	"fmt"
	"github.com/Webglhost-QA-Backend/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	_ "go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var MongoClient *mongo.Client

func InitMongo(cfg *config.MongoConfig) (*mongo.Client, error) {
	url := fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.USER, cfg.PWD, cfg.HOST, cfg.PORT)
	clientOptions := options.Client().ApplyURI(url)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		_ = client.Disconnect(ctx)
		return nil, err
	}
	log.Printf("Connected to MongoDB at %s", url)
	MongoClient = client
	return MongoClient, nil
}

func Close() error {
	if MongoClient != nil {
		return MongoClient.Disconnect(context.Background())
	}
	return nil
}
