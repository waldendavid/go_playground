package library

import "context"

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Library interface {
	GetBooks(ctx context.Context) ([]Book, error)
	// GetBook
	// CreateBook
	// UpdateBook
	// DeleteBook
}

func NewService() Library {
	return &service{}
}

type service struct {
}

func (s *service) GetBooks(ctx context.Context) ([]Book, error) {
	return nil, nil
}
