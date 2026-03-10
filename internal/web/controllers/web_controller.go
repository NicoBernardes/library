package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	bookService "library/internal/books/models"
	loanService "library/internal/loans/models"
	userService "library/internal/users/models"

	bookModel "library/internal/books/models"
	loanModel "library/internal/loans/models"
	userModel "library/internal/users/models"

	"github.com/gin-gonic/gin"
)

type WebController struct {
	templates   *template.Template
	bookService bookService.BookService
	userService userService.UserService
	loanService loanService.LoanService
}

func NewWebController(
	bookService bookService.BookService,
	userService userService.UserService,
	loanService loanService.LoanService,
) *WebController {
	tmpl := template.Must(template.ParseGlob("templates/*.html"))

	return &WebController{
		templates:   tmpl,
		bookService: bookService,
		userService: userService,
		loanService: loanService,
	}
}

func (wc *WebController) RegisterRoutes(router *gin.Engine) {
	router.GET("/", wc.ServeHome)
	router.GET("/users", wc.ServeUsers)
	router.GET("/books", wc.ServeBooks)
	router.GET("/loans", wc.ServeLoans)

	router.POST("/books", wc.CreateBook)
	router.POST("/users", wc.CreateUser)
	router.POST("/loans", wc.CreateLoan)

	router.GET("/loans/search", wc.SearchLoan)
	router.GET("/users/search", wc.SearchUser)
	router.GET("/books/search", wc.SearchBook)

	router.POST("/users/:id/edit", wc.UpdateUser)
	router.POST("/books/:id/edit", wc.UpdateBook)
	router.POST("/loans/:id/return", wc.ReturnBook)

	router.GET("/users/:id/edit", wc.EditUserForm)
	router.GET("/books/:id/edit", wc.EditBookForm)

	router.POST("/users/:id/delete", wc.DeleteUser)
	router.POST("/books/:id/delete", wc.DeleteBook)
}

func (wc *WebController) ServeBooks(c *gin.Context) {
	books, err := wc.bookService.GetAllBooks()

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]interface{}{
		"Title":         "Gerenciamento de Livros",
		"Books":         books,
		"ActiveSection": "books",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao renderizar o template: %v", err)
		return
	}
}

func (wc *WebController) SearchBook(c *gin.Context) {
	query := c.Query("q")

	var books []*bookModel.Book
	var err error

	if query != "" {
		allBooks, err := wc.bookService.GetAllBooks()
		if err != nil {
			wc.addFlashMessage(c, "Erro ao buscar livros: "+err.Error(), "error")
			c.Redirect(http.StatusSeeOther, "/books")
			return
		}

		for _, book := range allBooks {
			if contains(book.Title, query) || contains(book.Author, query) {
				books = append(books, book)
			}
		}
	} else {
		books, err = wc.bookService.GetAllBooks()
	}

	if err != nil {
		wc.addFlashMessage(c, "|Erro ao buscar livros: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]interface{}{
		"Title":         "Busca de Livros",
		"Books":         books,
		"ActiveSection": "books",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"SearchQuery":   query,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao renderizar template: %v", err)
		return
	}
}

func (wc *WebController) SearchUser(c *gin.Context) {
	query := c.Query("q")

	var users []*userModel.User
	var err error

	if query != "" {
		allUsers, err := wc.userService.GetAllUsers()
		if err != nil {
			wc.addFlashMessage(c, "Erro ao buscar usuários: "+err.Error(), "error")
			c.Redirect(http.StatusSeeOther, "/users")
			return
		}

		for _, user := range allUsers {
			if contains(user.Name, query) || contains(user.Email, query) {
				users = append(users, user)
			}
		}
	} else {
		users, err = wc.userService.GetAllUsers()
	}

	if err != nil {
		wc.addFlashMessage(c, "|Erro ao buscar usuários: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]interface{}{
		"Title":         "Busca de Usuários",
		"Users":         users,
		"ActiveSection": "users",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"SearchQuery":   query,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao renderizar template: %v", err)
		return
	}
}

func (wc *WebController) SearchLoan(c *gin.Context) {
	query := c.Query("q")
	status := c.Query("status")

	var loans []*loanModel.Loan
	var err error

	loans, err = wc.loanService.GetAllLoans()
	if err != nil {
		wc.addFlashMessage(c, "Erro ao buscar empréstimos: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/loans")
		return
	}

	if status != "" {
		var filteredLoans []*loanModel.Loan
		for _, loan := range loans {
			if loan.Status == status {
				filteredLoans = append(filteredLoans, loan)
			}
		}
		loans = filteredLoans
	}

	books, _ := wc.bookService.GetAllBooks()
	users, _ := wc.userService.GetAllUsers()

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]interface{}{
		"Title":         "Busca de Empréstimos",
		"Loans":         loans,
		"Books":         books,
		"Users":         users,
		"ActiveSection": "loans",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"SearchQuery":   query,
		"StatusFilter":  status,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao renderizar template: %v", err)
		return
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr ||
		(len(s) > len(substr) && (s[:len(substr)] == substr ||
			s[len(s)-len(substr):] == substr ||
			containsSubstring(s, substr))))
}

func containsSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}

	return false
}

func (wc *WebController) ReturnBook(c *gin.Context) {
	loanIDStr := c.Param("id")
	loanID, err := strconv.ParseInt(loanIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "ID do empréstimo inválido", "error")
		c.Redirect(http.StatusSeeOther, "/loans")
		return
	}

	err = wc.loanService.ReturnBook(loanID)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao devolver o livro: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(c, "Livro devolvido com sucesso!", "success")
	}

	c.Redirect(http.StatusSeeOther, "/loans")

}

func (wc *WebController) ServeLoans(c *gin.Context) {
	books, err := wc.bookService.GetAllBooks()
	users, err := wc.userService.GetAllUsers()
	loans, err := wc.loanService.GetAllLoans()

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]interface{}{
		"Title":         "Gerenciamento de Empréstimos",
		"Books":         books,
		"Users":         users,
		"Loans":         loans,
		"ActiveSection": "loans",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao renderizar o template: %v", err)
		return
	}
}

func (wc *WebController) CreateBook(c *gin.Context) {
	title := c.PostForm("title")
	author := c.PostForm("author")
	quantityStr := c.PostForm("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		wc.addFlashMessage(c, "Quantidade Inválida: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}

	book := &bookModel.Book{
		Title:    title,
		Author:   author,
		Quantity: quantity,
	}

	err = wc.bookService.CreateBook(book)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao criar o livro: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(c, "Livro criado com sucesso!", "success")
	}

	c.Redirect(http.StatusSeeOther, "/books")
}

func (wc *WebController) CreateLoan(c *gin.Context) {
	userIDStr := c.PostForm("user_id")
	bookIDStr := c.PostForm("book_id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "ID do usuário inválido", "error")
		c.Redirect(http.StatusSeeOther, "/loans")
		return
	}
	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "ID do livro inválido", "error")
		c.Redirect(http.StatusSeeOther, "/loans")
		return
	}

	_, err = wc.loanService.CreateLoan(bookID, userID)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao criar empréstimo: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(c, "empréstimo criado com sucesso!", "success")
	}

	c.Redirect(http.StatusSeeOther, "/loans")
}

func (wc *WebController) DeleteUser(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "ID do usuário inválido", "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	err = wc.userService.DeleteUser(userID)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao excluir o usuário: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) DeleteBook(c *gin.Context) {
	bookIDStr := c.Param("id")

	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "ID do livro inválido", "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}

	err = wc.bookService.DeleteBook(bookID)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao excluir o livro: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}

	c.Redirect(http.StatusSeeOther, "/books")
}

func (wc *WebController) UpdateUser(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "ID do usuário inválido", "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	user, err := wc.userService.GetUser(userID)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao buscar o usuario: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	name := c.PostForm("name")
	email := c.PostForm("email")
	user.Name = name
	user.Email = email

	err = wc.userService.UpdateUser(userID, user)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao atualizar o usuario: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) UpdateBook(c *gin.Context) {
	bookIDStr := c.Param("id")

	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "ID do livro inválido", "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}

	book, err := wc.bookService.GetBook(bookID)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao buscar o livro: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}

	title := c.PostForm("title")
	author := c.PostForm("author")
	quantityStr := c.PostForm("quantity")

	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		wc.addFlashMessage(c, "Quantidade inválida: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}

	book.Title = title
	book.Author = author
	book.Quantity = quantity

	err = wc.bookService.UpdateBook(bookID, book)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao atualizar o usuario: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}
	c.Redirect(http.StatusSeeOther, "/books")
}

func (wc *WebController) EditUserForm(c *gin.Context) {
	userIDStr := c.Param("id")

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "ID do usuário inválido", "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	user, err := wc.userService.GetUser(userID)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao buscar o usuario: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/users")
		return
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]interface{}{
		"Title":         "Editar Usuário",
		"User":          user,
		"ActiveSection": "users",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"IsEdit":        true,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao renderizar o template: %v", err)
		return
	}
}

func (wc *WebController) EditBookForm(c *gin.Context) {
	bookIDStr := c.Param("id")

	bookID, err := strconv.ParseInt(bookIDStr, 10, 64)
	if err != nil {
		wc.addFlashMessage(c, "ID do livro inválido", "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}

	book, err := wc.bookService.GetBook(bookID)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao buscar o usuario: "+err.Error(), "error")
		c.Redirect(http.StatusSeeOther, "/books")
		return
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]interface{}{
		"Title":         "Editar Livro",
		"Book":          book,
		"ActiveSection": "books",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"IsEdit":        true,
	}

	err = wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao renderizar o template: %v", err)
		return
	}
}

func (wc *WebController) ServeHome(c *gin.Context) {

	books, _ := wc.bookService.GetAllBooks()
	users, _ := wc.userService.GetAllUsers()
	loans, _ := wc.loanService.GetAllLoans()

	activeLoans := 0
	for _, loan := range loans {
		if loan.Status == "active" {
			activeLoans++
		}
	}

	avaiableBooks := 0
	for _, book := range books {
		if book.Quantity > 0 {
			avaiableBooks++
		}
	}

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]interface{}{
		"Title":         "Sistema de Biblioteca",
		"Books":         books,
		"Users":         users,
		"Loans":         loans,
		"ActiveSection": "dashboard",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
		"Stats": map[string]any{
			"TotalBooks":    len(books),
			"TotalUsers":    len(users),
			"TotalLoans":    len(loans),
			"ActiveLoans":   activeLoans,
			"AvaiableBooks": avaiableBooks,
		},
	}

	err := wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao renderizar o template: %v", err)
		return
	}

}

func (wc *WebController) ServeUsers(c *gin.Context) {
	users, _ := wc.userService.GetAllUsers()

	flashMessage, flashType := wc.getFlashMessage(c)

	data := map[string]interface{}{
		"Title":         "Gerenciamento de Usuários",
		"Users":         users,
		"ActiveSection": "users",
		"FlashMessage":  flashMessage,
		"FlashType":     flashType,
	}

	err := wc.templates.ExecuteTemplate(c.Writer, "layout", data)
	if err != nil {
		c.String(http.StatusInternalServerError, "Erro ao renderizar o template: %v", err)
		return
	}
}

func (wc *WebController) CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	user := &userModel.User{
		Name:  name,
		Email: email,
	}

	err := wc.userService.CreateUser(user)
	if err != nil {
		wc.addFlashMessage(c, "Erro ao criar o usuário: "+err.Error(), "error")
	} else {
		wc.addFlashMessage(c, "Usuário criado com sucesso!", "success")
	}

	c.Redirect(http.StatusSeeOther, "/users")
}

func (wc *WebController) addFlashMessage(c *gin.Context, message, messageType string) {
	c.SetCookie("flash_message", message, 1, "/", "", false, true)
	c.SetCookie("flash_type", messageType, 1, "/", "", false, true)
}

func (wc *WebController) getFlashMessage(c *gin.Context) (string, string) {
	message, _ := c.Cookie("flash_message")
	messageType, _ := c.Cookie("flash_type")

	c.SetCookie("flash_message", "", 1, "/", "", false, true)
	c.SetCookie("flash_type", "", 1, "/", "", false, true)

	return message, messageType
}
