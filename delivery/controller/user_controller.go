package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaqubmw/web-sales-app-golang/model"
	"github.com/yaqubmw/web-sales-app-golang/usecase"
)

type UserController struct {
	router *gin.Engine
	userUC usecase.UserUsecase
}

func (u *UserController) Register(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := u.userUC.Create(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userRes := map[string]any{
		"id":    user.Id,
		"nama":  user.Nama,
		"email": user.Email,
	}

	c.JSON(http.StatusCreated, userRes)
}

func NewUserController(router *gin.Engine, userUC usecase.UserUsecase) *UserController {
	controller := UserController{
		router: router,
		userUC: userUC,
	}

	rg := router.Group("/api/users")
	rg.POST("/register", controller.Register)

	return &controller
}
