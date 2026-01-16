package main

import (
	bookscontroller "library/internal/books/controllers"
	loancontroller "library/internal/loans/controllers"
	usercontroller "library/internal/users/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	booksController := bookscontroller.NewBooksController()
	userController := usercontroller.NewUserController()
	loanController := loancontroller.NewLoanController()

	booksController.RegisterRoutes(router)
	userController.RegisterRoutes(router)
	loanController.RegisterRoutes(router)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
