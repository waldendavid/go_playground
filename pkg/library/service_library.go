package library

import (
	"context"
	"fmt"

	//_ "sync"

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
	// tutaj get po tytule i sprawdzenie czy jest w bazie
	// res, err := s.olClient.Search(ctx, openlibrary.SearchRequest{Title: book.Title})
	// TODO zmienić kolejność, createBook na koniec wraz z zapisanie do cachea
	// if book.Title == "" {
	// }
	cb, inCache := s.cache.Get(book.Title)
	if inCache {
		return cb.(Book), nil
	}
	dbb, err := s.repo.GetBookByTitle(ctx, book.Title)
	if err == nil {
		s.cache.Set(book.Title, dbb)
		return dbb, nil
	}
	clires, err := s.olClient.Search(ctx, openlibrary.SearchRequest{Title: book.Title})
	if err == nil {
		//TODO book zgodnie z danymi z res
		var b Book
		b.Title = clires.Docs[len(clires.Docs)-1].Title
		s.cache.Set(book.Title, b)
		return s.repo.CreateBook(ctx, b)
	}
	return Book{}, fmt.Errorf("GetBook: %v", err)
}

func (s *service) UpdateBook(ctx context.Context, b Book, id string) (Book, error) {
	return s.repo.UpdateBook(ctx, b, id)
}

func (s *service) DeleteBook(ctx context.Context, id string) error {
	return s.repo.DeleteBook(ctx, id)
}
