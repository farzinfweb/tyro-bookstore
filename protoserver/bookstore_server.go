package protoserver

import (
	"bookstore/domain"
	"bookstore/protos"
	"context"
	"github.com/google/uuid"
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
	//TODO implement me
	panic("implement me")
}

func NewBookStoreServer(repo domain.IBookRepo) protos.BookstoreServer {
	return BookStoreServer{repo: repo}
}
