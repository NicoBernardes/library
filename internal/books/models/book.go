package models

import "time"

type Book struct {
	ID        int64     `json:"ID"`
	Title     string    `json:"title" binding:"required,min=5"`
	Author    string    `json:"author" binding:"required,min=5"`
	Quantity  int       `json:"quantity" binding:"required,min=1"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
