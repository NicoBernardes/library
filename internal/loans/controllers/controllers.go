package controllers

import "github.com/gin-gonic/gin"

type LoanController struct {
}

func NewLoanController() *LoanController {
	return &LoanController{}
}

func (c *LoanController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/loans")

	{
		users.POST("", c.CreateLoan)
		users.GET("/:id", c.GetLoan)
		users.GET("", c.GetAllLoans)
		users.PUT("/:id", c.UpdateLoan)
		users.DELETE("/:id", c.DeleteLoan)
	}
}

func (c *LoanController) CreateLoan(ctx *gin.Context) {

}

func (c *LoanController) GetLoan(ctx *gin.Context) {

}

func (c *LoanController) GetAllLoans(ctx *gin.Context) {

}

func (c *LoanController) UpdateLoan(ctx *gin.Context) {

}

func (c *LoanController) DeleteLoan(ctx *gin.Context) {

}
