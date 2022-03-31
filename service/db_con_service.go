package service

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectionInfo is connection data info
type ConnectionInfo struct {
	DBName         string
	CollectionName string
}

// MongoDbContext is implemet db connect action
func MongoDbContext(c ConnectionInfo) *mongo.Collection {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://j_dev:" + os.Getenv("MONGODBPSWD") + "@jdev.y4x5s.gcp.mongodb.net/" + c.DBName + "?retryWrites=true&w=majority").
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// 使用完關閉
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 嘗試連線
	// 錯誤處理
	if err != nil {
		log.Fatal(err)
	}

	return (*mongo.Collection)(client.Database(c.DBName).Collection(c.CollectionName))
}
