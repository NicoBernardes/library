package controllers

import (
	"library/internal/books/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BooksController struct {
	bookService models.BookService
}

func NewBooksController(bookService models.BookService) *BooksController {
	return &BooksController{bookService: bookService}
}

func (b *BooksController) RegisterRoutes(r *gin.RouterGroup) {
	users := r.Group("/books")

	{
		users.POST("", b.CreateBook)
		users.GET("/:id", b.GetBook)
		users.GET("", b.GetAllBooks)
		users.PUT("/:id", b.UpdateBook)
		users.DELETE("/:id", b.DeleteBook)
	}
}

func (b *BooksController) CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := b.bookService.CreateBook(&book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, book)
}

func (b *BooksController) GetBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := b.bookService.GetBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (b *BooksController) GetAllBooks(ctx *gin.Context) {
	book, err := b.bookService.GetAllBooks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, book)
}

func (b *BooksController) UpdateBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	err = b.bookService.UpdateBook(id, &book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}

func (b *BooksController) DeleteBook(ctx *gin.Context) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = b.bookService.DeleteBook(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, nil)
}
