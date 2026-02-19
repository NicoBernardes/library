package main

import (
	bookRepository "library/internal/books/repositories"
	loanRepository "library/internal/loans/repositories"
	userRepository "library/internal/users/repositories"

	bookService "library/internal/books/services"
	loanService "library/internal/loans/services"
	userService "library/internal/users/services"

	booksController "library/internal/books/controllers"
	loanController "library/internal/loans/controllers"
	userController "library/internal/users/controllers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	bookRepository := bookRepository.NewBookRepository()
	loanRepository := loanRepository.NewLoanRepository()
	usersRepository := userRepository.NewUserRepository()

	bookService := bookService.NewBookService(bookRepository)
	userService := userService.NewUserService(usersRepository)
	loanService := loanService.NewLoanService(loanRepository, bookService, userService)

	booksController := booksController.NewBooksController(bookService)
	userController := userController.NewUserController(userService)
	loanController := loanController.NewLoanController(loanService)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	booksController.RegisterRoutes(router)
	userController.RegisterRoutes(router)
	loanController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
