package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) RegisterRoutes(r *gin.Engine) {
	users := r.Group("/users")

	{
		users.POST("", c.CreateUser)
		users.GET("/:id", c.GetUser)
		users.GET("", c.GetAllUsers)
		users.PUT("/:id", c.UpdateUser)
		users.DELETE("/:id", c.DeleteUser)
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {

}

func (c *UserController) GetUser(ctx *gin.Context) {

}

func (c *UserController) GetAllUsers(ctx *gin.Context) {

}

func (c *UserController) UpdateUser(ctx *gin.Context) {

}

func (c *UserController) DeleteUser(ctx *gin.Context) {

}
