package repositories

import (
	"errors"
	"library/internal/books/models"
	"sync"
)

type BookRepository struct {
	books  map[int64]*models.Book
	mu     sync.RWMutex
	nextId int64
}

func NewBookRepository() models.BookRepository {
	return &BookRepository{
		books:  make(map[int64]*models.Book),
		nextId: 1,
	}
}

// CreateBook implements [models.BookRepository].
func (b *BookRepository) CreateBook(book *models.Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	book.ID = b.nextId
	b.nextId++
	b.books[book.ID] = book
	return nil
}

// DeleteBook implements [models.BookRepository].
func (b *BookRepository) DeleteBook(id int64) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	_, exists := b.books[id]
	if !exists {
		return errors.New("book not found")
	}

	delete(b.books, id)
	return nil
}

// GetAllBooks implements [models.BookRepository].
func (b *BookRepository) GetAllBooks() ([]*models.Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	books := make([]*models.Book, 0, len(b.books))
	for _, book := range b.books {
		books = append(books, book)
	}

	return books, nil
}

// GetBook implements [models.BookRepository].
func (b *BookRepository) GetBook(id int64) (*models.Book, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	book, exists := b.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}

	return book, nil
}

// UpdateBook implements [models.BookRepository].
func (b *BookRepository) UpdateBook(id int64, book *models.Book) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	book, exists := b.books[id]
	if !exists {
		return errors.New("book not found")
	}
	//id declarado diretamente pois será enviado no parametro para update
	b.books[id] = book

	return nil
}
