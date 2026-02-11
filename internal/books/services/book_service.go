package services

import (
	"errors"
	"library/internal/books/models"
)

type BookService struct {
	bookRepository models.BookRepository
}

func NewBookService(bookRepository models.BookRepository) models.BookService {
	return &BookService{bookRepository: bookRepository}
}

func (b *BookService) CreateBook(book *models.Book) error {

	if book.Quantity < 1 {
		return errors.New("quantity cannot be under 0")
	}

	return b.bookRepository.CreateBook(book)
}

func (b *BookService) GetBook(id int64) (*models.Book, error) {
	return b.bookRepository.GetBook(id)
}

func (b *BookService) GetAllBooks() ([]*models.Book, error) {
	return b.bookRepository.GetAllBooks()
}

func (b *BookService) UpdateBook(id int64, book *models.Book) error {
	return b.bookRepository.UpdateBook(id, book)
}

func (b *BookService) DeleteBook(id int64) error {
	return b.bookRepository.DeleteBook(id)
}
