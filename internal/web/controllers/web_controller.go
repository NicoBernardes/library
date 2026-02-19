package controllers

import (
	"html/template"
	"net/http"

	bookService "library/internal/books/models"
	loanService "library/internal/loans/models"
	userService "library/internal/users/models"

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

	router.POST("/users", wc.CreateUser)
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
