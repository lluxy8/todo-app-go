package main

import (
	"context"
	"log"
	"time"

	"github.com/lluxy8/todo-app-go/internal/config"
	"github.com/lluxy8/todo-app-go/internal/router"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongo := connectMongo(cfg)
	defer func ()  {
		_= mongo.Disconnect(context.Background())
	}()

	log.Println("MongoDB connected successfully")
	log.Println("App running on port:", cfg.App.Port)

	r := router.New(mongo)
	if err := r.Run(":" + cfg.App.Port); err != nil {
		log.Fatal(err)
	}
}

func connectMongo(cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Mongo.URI()))
	if err != nil {
		log.Fatalf("mongo connect error: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("mongo ping error: %v", err)
	}

	return client
}
