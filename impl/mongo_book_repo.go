package impl

import (
	"bookstore/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoBook struct {
	Id        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Price     uint32             `bson:"price"`
	Author    string             `bson:"author"`
	CreatedAt time.Time          `bson:"createdAt"`
}

type MongoBookRepo struct {
	collection *mongo.Collection
}

func (m MongoBookRepo) GetById(ctx context.Context, id string) (domain.Book, error) {
	oid, _ := primitive.ObjectIDFromHex(id)
	var book MongoBook
	err := m.collection.FindOne(ctx, bson.M{"_id": oid}).Decode(&book)
	if err != nil {
		return domain.Book{}, err
	}
	return domain.Book{
		Id:        book.Id.Hex(),
		Title:     book.Title,
		Author:    book.Author,
		Price:     book.Price,
		CreatedAt: book.CreatedAt,
	}, nil
}

func (m MongoBookRepo) ListAll(ctx context.Context, params domain.ListBookStoreParams) ([]domain.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewMongoBookRepo(collection *mongo.Collection) domain.IBookRepo {
	return MongoBookRepo{collection}
}
