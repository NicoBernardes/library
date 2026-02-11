package models

type BookRepository interface {
	CreateBook(book *Book) error
	GetBook(id int64) (*Book, error)
	GetAllBooks() ([]*Book, error)
	UpdateBook(id int64, book *Book) error
	DeleteBook(id int64) error
}
