package impl

import (
	"bookstore/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	return m.toDomainBook(book), nil
}

func (m MongoBookRepo) ListAll(ctx context.Context, params domain.ListBookStoreParams) ([]domain.Book, uint32, error) {
	filters := bson.D{}
	if params.SearchTerm != "" {
		filters = append(filters, bson.E{Key: "title", Value: bson.D{{"$regex", primitive.Regex{Pattern: params.SearchTerm, Options: "i"}}}})
	}
	limit := int64(params.PerPage)
	skip := int64(params.PerPage * (params.Page - 1))
	opts := &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	}
	cursor, err := m.collection.Find(ctx, filters, opts)
	if err != nil {
		return nil, 0, err
	}
	count, err := m.collection.CountDocuments(ctx, filters)
	if err != nil {
		return nil, 0, err
	}
	var books []MongoBook
	if err := cursor.All(ctx, &books); err != nil {
		return nil, 0, err
	}
	return m.toDomainBooks(books), uint32(count), nil
}

func (m MongoBookRepo) toDomainBook(book MongoBook) domain.Book {
	return domain.Book{
		Id:        book.Id.Hex(),
		Title:     book.Title,
		Author:    book.Author,
		Price:     book.Price,
		CreatedAt: book.CreatedAt,
	}
}

func (m MongoBookRepo) toDomainBooks(books []MongoBook) []domain.Book {
	var _books []domain.Book
	for _, b := range books {
		_books = append(_books, m.toDomainBook(b))
	}
	return _books
}

func NewMongoBookRepo(collection *mongo.Collection) domain.IBookRepo {
	return MongoBookRepo{collection}
}
