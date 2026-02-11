package models

import "time"

type Loan struct {
	ID         int64     `json:"ID"`
	BookID     int64     `json:"bookID"`
	UserID     int64     `json:"userID"`
	BorrowedAt time.Time `json:"borrowedAt"`
	ReturnedAt time.Time `json:"returnedAt"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
