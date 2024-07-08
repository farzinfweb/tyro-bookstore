package protoserver

import (
	"bookstore/domain"
	"bookstore/protos"
	"context"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BookStoreServer struct {
	repo domain.IBookRepo
	protos.UnimplementedBookstoreServer
}

func (b BookStoreServer) Buy(ctx context.Context, req *protos.BuyReq) (*protos.BuyResp, error) {
	book, err := b.repo.GetById(ctx, req.BookId)
	if err != nil {
		return nil, err
	}
	orderId, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	return &protos.BuyResp{
		Status:  0,
		OrderId: orderId.String(),
		Price:   req.Quantity * book.Price,
	}, nil
}

func (b BookStoreServer) Search(ctx context.Context, req *protos.SearchReq) (*protos.SearchResp, error) {
	books, count, err := b.repo.ListAll(ctx, domain.ListBookStoreParams{
		SearchTerm: req.SearchTerm,
		Page:       req.Page,
		PerPage:    req.PerPage,
	})
	if err != nil {
		return nil, err
	}
	var _books = make([]*protos.Book, len(books))
	for i := 0; i < len(books); i++ {
		_b := books[i]
		_books[i] = &protos.Book{
			Id:        _b.Id,
			Title:     _b.Title,
			Author:    _b.Author,
			Price:     _b.Price,
			CreatedAt: timestamppb.New(_b.CreatedAt),
		}
	}
	return &protos.SearchResp{
		Result:     _books,
		TotalCount: count,
	}, nil
}

func NewBookStoreServer(repo domain.IBookRepo) protos.BookstoreServer {
	return BookStoreServer{repo: repo}
}
