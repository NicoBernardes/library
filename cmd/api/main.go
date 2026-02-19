package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	bookRepositoryPkg "library/internal/books/repositories"
	loanRepositoryPkg "library/internal/loans/repositories"
	userRepositoryPkg "library/internal/users/repositories"

	bookServicePkg "library/internal/books/services"
	loanServicePkg "library/internal/loans/services"
	userServicePkg "library/internal/users/services"

	booksControllerPkg "library/internal/books/controllers"
	loanControllerPkg "library/internal/loans/controllers"
	userControllerPkg "library/internal/users/controllers"
	webControllerPkg "library/internal/web/controllers"
)

func main() {
	router := gin.Default()

	bookRepository := bookRepositoryPkg.NewBookRepository()
	loanRepository := loanRepositoryPkg.NewLoanRepository()
	usersRepository := userRepositoryPkg.NewUserRepository()

	bookService := bookServicePkg.NewBookService(bookRepository)
	userService := userServicePkg.NewUserService(usersRepository)
	loanService := loanServicePkg.NewLoanService(loanRepository, bookService, userService)

	booksController := booksControllerPkg.NewBooksController(bookService)
	userController := userControllerPkg.NewUserController(userService)
	loanController := loanControllerPkg.NewLoanController(loanService)
	webController := webControllerPkg.NewWebController(bookService, userService, loanService)

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	router.Use(cors.New(config))

	apiGroup := router.Group("/api")

	booksController.RegisterRoutes(apiGroup)
	userController.RegisterRoutes(apiGroup)
	loanController.RegisterRoutes(apiGroup)

	webController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
