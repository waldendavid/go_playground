package library

import (
	"context"
	"errors"

	"github.com/pborman/uuid"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Library interface {
	GetBooks(ctx context.Context) ([]Book, error)
	GetBook(ctx context.Context, id string) (Book, error)
	CreateBook(ctx context.Context, book Book) (Book, error)
	UpdateBook(ctx context.Context, book Book, id string) (Book, error)
	//Co zwraca delete? https://stackoverflow.com/questions/25970523/restful-what-should-a-delete-response-body-contain
	//204? W kodzie nic?
	DeleteBook(ctx context.Context, id string) error
}

func NewService() Library {
	var books []Book
	books = append(books, Book{ID: "1", Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	return &service{books: books}
}

type service struct {
	books []Book
}

//jak sprawdzić usage?
func (s *service) GetBooks(ctx context.Context) ([]Book, error) {
	return s.books, nil
}

func (s *service) GetBook(ctx context.Context, id string) (Book, error) {
	for _, b := range s.books {
		if b.ID == id {
			return b, nil
		}
	}

	return Book{}, errors.New("Book not found")
}

func (s *service) CreateBook(ctx context.Context, book Book) (Book, error) {
	book.ID = uuid.New()
	s.books = append(s.books, book)
	return book, nil
}

// TODO add errors
func (s *service) UpdateBook(ctx context.Context, book Book, id string) (Book, error) {
	for i := range s.books {
		if s.books[i].ID == id {
			book.ID = id
			s.books[i] = book
		}
	}
	return book, nil
}

func (s *service) DeleteBook(ctx context.Context, id string) error {
	var books []Book
	for _, b := range s.books {
		if b.ID == id {
			continue
		}
		books = append(books, b)
	}
	s.books = books
	return nil
}

//https://tutorialedge.net/golang/creating-restful-api-with-golang/
