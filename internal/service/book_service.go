package service

import (
	"context"
	"errors"
	"strings"

	"github.com/leeozebra/go-crud/internal/domain"
)

type BookRepository interface {
	Create(ctx context.Context, in domain.CreateBookInput) (domain.Book, error)
	GetByID(ctx context.Context, id string) (domain.Book, error)
	List(ctx context.Context, limit, offset int) ([]domain.Book, error)
	UpdatePartial(ctx context.Context, id string, in domain.UpdateBookInput) (domain.Book, error)
	Delete(ctx context.Context, id string) error
}

type BookService struct {
	repo BookRepository
}

func NewBookService(r BookRepository) *BookService { return &BookService{repo: r} }

func (s *BookService) Create(ctx context.Context, in domain.CreateBookInput) (domain.Book, error) {
	if strings.TrimSpace(in.Title) == "" || strings.TrimSpace(in.Author) == "" {
		return domain.Book{}, errors.New("title and author are required")
	}
	if in.Price < 0 {
		return domain.Book{}, errors.New("price must be >= 0")
	}
	return s.repo.Create(ctx, in)
}

func (s *BookService) GetByID(ctx context.Context, id string) (domain.Book, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *BookService) List(ctx context.Context, limit, offset int) ([]domain.Book, error) {
	return s.repo.List(ctx, limit, offset)
}

func (s *BookService) UpdatePartial(ctx context.Context, id string, in domain.UpdateBookInput) (domain.Book, error) {
	return s.repo.UpdatePartial(ctx, id, in)
}

func (s *BookService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
