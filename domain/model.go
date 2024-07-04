package domain

import "time"

type Book struct {
	Id        string
	Title     string
	Author    string
	Price     uint32
	CreatedAt time.Time
}

type ListBookStoreParams struct {
	SearchTerm string
	Page       uint32
	PerPage    uint32
}
