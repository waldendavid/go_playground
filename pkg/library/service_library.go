package library

import (
	"context"

	_ "sync"

	"github.com/waldendavid/restapi/pkg/cache"
	"github.com/waldendavid/restapi/pkg/openlibrary"
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

func NewService(repo Repository, olClient openlibrary.Client, cache cache.MemoryCache) Service {
	return &service{
		repo:     repo,
		olClient: olClient,
		cache:    cache,
	}
}

type service struct {
	repo     Repository
	olClient openlibrary.Client
	cache    cache.MemoryCache
}

func (s *service) GetBooks(ctx context.Context) ([]Book, error) {
	return s.repo.GetBooks(ctx)
}

func (s *service) GetBook(ctx context.Context, id string) (Book, error) {
	return s.repo.GetBook(ctx, id)
}

func (s *service) CreateBook(ctx context.Context, book Book) (Book, error) {
	// todo tutaj get po tytule i sprawdzenie czy jest w bazie
	// res, err := s.olClient.Search(ctx, openlibrary.SearchRequest{Title: book.Title})
	// if book.Title == "" {

	// }
	return s.repo.CreateBook(ctx, book)
}

func (s *service) UpdateBook(ctx context.Context, b Book, id string) (Book, error) {
	return s.repo.UpdateBook(ctx, b, id)
}

func (s *service) DeleteBook(ctx context.Context, id string) error {
	return s.repo.DeleteBook(ctx, id)
}
