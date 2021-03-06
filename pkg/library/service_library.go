package library

import (
	"context"
	"fmt"

	"github.com/waldendavid/restapi/pkg/cache"
	"github.com/waldendavid/restapi/pkg/openlibrary"
	"gorm.io/gorm"
)

// TODO
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

func NewService(repo Repository, olClient openlibrary.Client, cache cache.Cache) Service {
	return &service{
		repo:     repo,
		olClient: olClient,
		cache:    cache,
	}
}

type service struct {
	repo     Repository
	olClient openlibrary.Client
	cache    cache.Cache
}

func (s *service) GetBooks(ctx context.Context) ([]Book, error) {
	return s.repo.GetBooks(ctx)
}

func (s *service) GetBook(ctx context.Context, id string) (Book, error) {

	return s.repo.GetBook(ctx, id)
}

func (s *service) CreateBook(ctx context.Context, book Book) (Book, error) {
	cb, inCache := s.cache.Get(book.Title)
	if inCache {
		fmt.Println("In Cache")
		return cb.(Book), nil
	}
	clires, err := s.olClient.Search(ctx, openlibrary.SearchRequest{Title: book.Title})
	if err != nil {
		return Book{}, fmt.Errorf("GetBook: %v", err)

	}
	//TODO book zgodnie z danymi z res
	var b Book
	t := clires.Docs[len(clires.Docs)-1].Title
	b.Title = t
	s.cache.Set(book.Title, b)
	return s.repo.CreateBook(ctx, b)
}

func (s *service) UpdateBook(ctx context.Context, b Book, id string) (Book, error) {
	return s.repo.UpdateBook(ctx, b, id)
}

func (s *service) DeleteBook(ctx context.Context, id string) error {
	return s.repo.DeleteBook(ctx, id)
}
