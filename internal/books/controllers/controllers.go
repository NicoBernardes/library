package controllers

import "github.com/gin-gonic/gin"

type BooksController struct {
}

func NewBooksController() *BooksController {
	return &BooksController{}
}

func (c *BooksController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/books")

	{
		users.POST("", c.CreateBook)
		users.GET("/:id", c.GetBook)
		users.GET("", c.GetAllBooks)
		users.PUT("/:id", c.UpdateBook)
		users.DELETE("/:id", c.DeleteBook)
	}
}

func (c *BooksController) CreateBook(ctx *gin.Context) {

}

func (c *BooksController) GetBook(ctx *gin.Context) {

}

func (c *BooksController) GetAllBooks(ctx *gin.Context) {

}

func (c *BooksController) UpdateBook(ctx *gin.Context) {

}

func (c *BooksController) DeleteBook(ctx *gin.Context) {

}
