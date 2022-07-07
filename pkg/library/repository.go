package library

import "context"

type Repository interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, id string) (Book, error)
	CreateBook(ctx context.Context, book Book) (Book, error)
	UpdateBook(ctx context.Context, book Book, id string) (Book, error)
	DeleteBook(ctx context.Context, id string) error
}
