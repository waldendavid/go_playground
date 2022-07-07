package library

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewRepositoryGorm() Repository {
	dsn := "host=localhost user=postgres password=secret dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Book{})
	db.Create(&Book{Isbn: "44778854", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	db.Create(&Book{Isbn: "3987654", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})
	return &repository{db: db}
}

type repository struct {
	db *gorm.DB
}

func (repo *repository) GetBooks(ctx context.Context) ([]Book, error) {
	var books []Book
	result := repo.db.Find(&books)
	if result.Error != nil {

		return nil, fmt.Errorf("GetBooks: %v", result.Error)
	}
	return books, nil
}

func (repo *repository) GetBook(ctx context.Context, id string) (Book, error) {
	book := Book{}
	result := repo.db.First(&book, id)
	if result.Error != nil {
		return Book{}, fmt.Errorf("GetBook: %v", result.Error)
	}

	return book, nil
}

func (repo *repository) CreateBook(ctx context.Context, book Book) (Book, error) {
	result := repo.db.Create(&book)
	if result.Error != nil {

		return Book{}, fmt.Errorf("CreateBook: %v", result.Error)
	}
	return book, nil
}

func (repo *repository) UpdateBook(ctx context.Context, b Book, id string) (Book, error) {
	book := Book{}
	result := repo.db.First(&book, id)
	if result.Error != nil {

		return Book{}, fmt.Errorf("UpdateBook: %v", result.Error)
	}

	book.Isbn = b.Isbn
	book.Title = b.Title
	//todo itd....
	result = repo.db.Save(&book) //variable shadowing
	if result.Error != nil {

		return Book{}, fmt.Errorf("UpdateBook: %v", result.Error)
	}
	return book, nil
}

func (repo *repository) DeleteBook(ctx context.Context, id string) error {
	result := repo.db.Delete(&Book{}, id)
	if result.Error != nil {

		return fmt.Errorf("DeleteBook: %v", result.Error)
	}
	return nil
}
