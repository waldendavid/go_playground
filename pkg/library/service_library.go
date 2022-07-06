package library

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

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

func NewService(db *gorm.DB) Library {
	// Migrate the schema: Author and Book
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
	db.Create(&Book{Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	db.Create(&Book{Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	return &service{db: db}
}

type service struct {
	db *gorm.DB
}

func (s *service) GetBooks(ctx context.Context) ([]Book, error) {
	var books []Book
	result := s.db.Find(&books)
	if result.Error != nil {

		return nil, fmt.Errorf("GetBooks: %v", result.Error)
	}
	return books, nil
}

func (s *service) GetBook(ctx context.Context, id string) (Book, error) {
	book := Book{}
	result := s.db.First(&book, id)
	if result.Error != nil {
		return Book{}, fmt.Errorf("GetBook: %v", result.Error)
	}

	return book, nil
}

func (s *service) CreateBook(ctx context.Context, book Book) (Book, error) {
	result := s.db.Create(&book)
	if result.Error != nil {

		return Book{}, fmt.Errorf("CreateBook: %v", result.Error)
	}
	return book, nil
}

func (s *service) UpdateBook(ctx context.Context, b Book, id string) (Book, error) {
	book := Book{}
	result := s.db.First(&book, id)
	if result.Error != nil {

		return Book{}, fmt.Errorf("UpdateBook: %v", result.Error)
	}

	book.Isbn = b.Isbn
	book.Title = b.Title
	//todo itd....
	result = s.db.Save(&book) //variable shadowing
	if result.Error != nil {

		return Book{}, fmt.Errorf("UpdateBook: %v", result.Error)
	}
	return book, nil
}

func (s *service) DeleteBook(ctx context.Context, id string) error {
	result := s.db.Delete(&Book{}, id)
	if result.Error != nil {

		return fmt.Errorf("DeleteBook: %v", result.Error)
	}
	return nil
}
