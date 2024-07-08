package domain

import "context"

type IBookRepo interface {
	GetById(ctx context.Context, id string) (Book, error)
	ListAll(ctx context.Context, params ListBookStoreParams) ([]Book, uint32, error)
}
