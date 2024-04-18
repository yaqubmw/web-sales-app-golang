package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yaqubmw/web-sales-app-golang/delivery/middleware"
	"github.com/yaqubmw/web-sales-app-golang/model"
	"github.com/yaqubmw/web-sales-app-golang/usecase"
)

type AuthController struct {
	router *gin.Engine
	authUC usecase.AuthUsecase
}

func (a *AuthController) Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	token, err := a.authUC.Login(user.Email, user.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (a *AuthController) Logout(c *gin.Context) {
	var token model.Token
	if err := c.ShouldBindHeader(&token); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := a.authUC.Logout(token.UserId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logout success"})
}

func NewAuthController(router *gin.Engine, authUC usecase.AuthUsecase) *AuthController {
	controller := AuthController{
		router: router,
		authUC: authUC,
	}

	rg := router.Group("/api/auth")
	rg.POST("/login", controller.Login)
	rg.POST("/logout", middleware.AuthMiddleware(), controller.Logout)

	return &controller
}