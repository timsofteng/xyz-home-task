package service

import (
	"context"
)

// RepoInterface represents all repo handlers.
type RepoInterface interface {
	// Get all books
	GetBooks(ctx context.Context, query string) ([]Book, error)
}

type BooksService struct{ repo RepoInterface }

func New(repo RepoInterface) *BooksService {
	return &BooksService{
		repo: repo,
	}
}

func (s *BooksService) GetBooks(ctx context.Context, query string) ([]Book, error) {
	books, err := s.repo.GetBooks(ctx, query)
	return books, err
}
