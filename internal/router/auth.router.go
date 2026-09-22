package router

import (
	"github.com/gin-gonic/gin"
	"github.com/habibmrizki/week10/internal/handler"
	"github.com/habibmrizki/week10/internal/service"
)

func InitAuthRouter(router *gin.Engine) {
	usersGroup := router.Group("/auth")

	authService := service.NewAuthService()
	authHandler := handler.NewAuthHandler(authService)

	usersGroup.POST("/register", authHandler.Register)
	usersGroup.POST("/login", authHandler.Login)
}
