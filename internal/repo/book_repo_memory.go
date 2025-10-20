package repo

import (
	"context"
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/leeozebra/go-crud/internal/domain"
)

var ErrNotFound = errors.New("not found")

type BookRepoMem struct {
	mu    sync.RWMutex
	store map[string]domain.Book
}

func NewBookRepoMem() *BookRepoMem {
	return &BookRepoMem{store: make(map[string]domain.Book)}
}

func (r *BookRepoMem) Create(ctx context.Context, in domain.CreateBookInput) (domain.Book, error) {
	now := time.Now().UTC()
	b := domain.Book{
		ID:        uuid.New().String(),
		Title:     in.Title,
		Author:    in.Author,
		Price:     in.Price,
		CreatedAt: now,
	}
	r.mu.Lock()
	r.store[b.ID] = b
	r.mu.Unlock()
	return b, nil
}

func (r *BookRepoMem) GetByID(ctx context.Context, id string) (domain.Book, error) {
	r.mu.RLock()
	b, ok := r.store[id]
	r.mu.RUnlock()
	if !ok {
		return domain.Book{}, ErrNotFound
	}
	return b, nil
}

func (r *BookRepoMem) List(ctx context.Context, limit, offset int) ([]domain.Book, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	all := make([]domain.Book, 0, len(r.store))
	for _, v := range r.store {
		all = append(all, v)
	}

	sort.Slice(all, func(i, j int) bool { return all[i].CreatedAt.After(all[j].CreatedAt) })

	if limit <= 0 || limit > 200 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	end := offset + limit
	if offset > len(all) {
		return []domain.Book{}, nil
	}
	if end > len(all) {
		end = len(all)
	}
	return all[offset:end], nil
}

func (r *BookRepoMem) UpdatePartial(ctx context.Context, id string, in domain.UpdateBookInput) (domain.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	b, ok := r.store[id]
	if !ok {
		return domain.Book{}, ErrNotFound
	}
	if in.Title != nil {
		b.Title = *in.Title
	}
	if in.Author != nil {
		b.Author = *in.Author
	}
	if in.Price != nil {
		b.Price = *in.Price
	}
	now := time.Now().UTC()
	b.UpdatedAt = &now
	r.store[id] = b
	return b, nil
}

func (r *BookRepoMem) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.store[id]; !ok {
		return ErrNotFound
	}
	delete(r.store, id)
	return nil
}
