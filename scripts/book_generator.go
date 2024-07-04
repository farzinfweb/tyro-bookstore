package main

import (
	"bookstore/impl"
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("bookstore").Collection("books")

	books := make([]interface{}, 100)
	for i := 0; i < 100; i++ {
		books[i] = impl.MongoBook{
			Id:        primitive.NewObjectID(),
			Title:     gofakeit.BookTitle(),
			Author:    gofakeit.BookAuthor(),
			Price:     uint32(gofakeit.Price(500000, 5000000)),
			CreatedAt: time.Now(),
		}
	}

	_, err = collection.InsertMany(context.TODO(), books)
	if err != nil {
		return
	}

	log.Info("100 books created")
}
