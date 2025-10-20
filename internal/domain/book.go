package domain

import "time"

type Book struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Author    string     `json:"author"`
	Price     float64    `json:"price"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}

type CreateBookInput struct {
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

type UpdateBookInput struct {
	Title  *string  `json:"title,omitempty"`
	Author *string  `json:"author,omitempty"`
	Price  *float64 `json:"price,omitempty"`
}
