package library

import "context"

type Repository interface {
	//TODO check what should be put in arg
	GetBooks(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, id string) (Book, error)
	CreateBook(ctx context.Context, book Book) (Book, error)
	UpdateBook(ctx context.Context, book Book, id string) (Book, error)
	DeleteBook(ctx context.Context, id string) error
}
