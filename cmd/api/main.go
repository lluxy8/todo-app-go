// @title Todo API
// @version 1.0.0
// @description This is a simple Todo API server written in Go using MongoDB.
// @contact.name Murat Can Sahin
// @contact.email lluxysa@gmail.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @host localhost:8080
// @BasePath /

package main

import (
	"context"
	"log"
	"time"

	"github.com/lluxy8/todo-app-go/internal/config"
	"github.com/lluxy8/todo-app-go/internal/router"

	_ "github.com/lluxy8/todo-app-go/docs"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongo := connectMongo(cfg)
	defer func() {
		_ = mongo.Disconnect(context.Background())
	}()

	log.Println("MongoDB connected successfully")

	db := mongo.Database(cfg.Mongo.Database)
	todoCollection := db.Collection("todos")
	r := router.New(router.RouterDeps{
		Cfg:            cfg,
		TodoCollection: todoCollection,
	})
	if err := r.Run(":" + cfg.App.Port); err != nil {
		log.Fatal(err)
	}

	log.Println("App running on port:", cfg.App.Port)
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
