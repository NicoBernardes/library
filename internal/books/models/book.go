package models

import "time"

type Book struct {
	ID        int64     `json:"ID"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
