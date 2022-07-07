package library

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

// todo
type BookDTO struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Book struct {
	gorm.Model
	Isbn   string
	Title  string
	Author *Author
}

type AuthorDTO struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Author struct {
	gorm.Model
	Firstname string
	Lastname  string
	BookID    uint
}

func NewService(repo Repository) Service {
	// Migrate the schema: Author and Book
	// db.AutoMigrate(&Author{})
	// db.AutoMigrate(&Book{})
	// db.Create(&Book{Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	// db.Create(&Book{Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	return &service{repo: repo}
}

type service struct {
	repo Repository
}

func (s *service) GetBooks(ctx context.Context) ([]Book, error) {
	var books []Book
	// todo should repo.getBooks have books, err? err is checked in repository
	books, err := s.repo.GetBooks(ctx)
	if err != nil {

		return nil, fmt.Errorf("GetBooks: %v", err)
	}
	return books, nil
}

func (s *service) GetBook(ctx context.Context, id string) (Book, error) {
	book := Book{}
	book, err := s.repo.GetBook(ctx, id)
	if err != nil {
		return Book{}, fmt.Errorf("GetBook: %v", err)
	}

	return book, nil
}

func (s *service) CreateBook(ctx context.Context, book Book) (Book, error) {
	book, err := s.repo.CreateBook(ctx, book)
	if err != nil {

		return Book{}, fmt.Errorf("CreateBook: %v", err)
	}
	return book, nil
}

func (s *service) UpdateBook(ctx context.Context, b Book, id string) (Book, error) {
	book := Book{}
	book, err := s.repo.UpdateBook(ctx, b, id)
	if err != nil {

		return Book{}, fmt.Errorf("UpdateBook: %v", err)
	}

	// book.Isbn = b.Isbn
	// book.Title = b.Title
	// //todo itd....
	// book = s.repo.(&book) //variable shadowing
	// if result.Error != nil {

	// 	return Book{}, fmt.Errorf("UpdateBook: %v", result.Error)
	// }
	return book, nil
}

func (s *service) DeleteBook(ctx context.Context, id string) error {
	// result :=
	s.repo.DeleteBook(ctx, id)
	// if result.Error != nil {

	// 	return fmt.Errorf("DeleteBook: %v", result.Error)
	// }
	return nil
}
