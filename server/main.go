package main

import (
	"bookstore/impl"
	"bookstore/protos"
	"bookstore/protoserver"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"net"
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

	listen, err := net.Listen("tcp", "localhost:7080")
	if err != nil {
		panic(err)
	}

	db := client.Database("bookstore")

	repo := impl.NewMongoBookRepo(db.Collection("books"))

	srv := protoserver.NewBookStoreServer(repo)

	s := grpc.NewServer()
	protos.RegisterBookstoreServer(s, srv)

	err = s.Serve(listen)
	if err != nil {
		panic(err)
	}
}
