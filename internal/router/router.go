package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/habibmrizki/week10/internal/dto"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	InitAuthRouter(router)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, dto.Response{
			Msg:     "rute salah",
			Success: false,
		})
	})

	return router
}
