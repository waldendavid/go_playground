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
	var books []Book
	books = append(books, Book{ID: "1", Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	return books, nil
}
