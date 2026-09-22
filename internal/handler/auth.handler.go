package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/habibmrizki/week10/internal/dto"
	"github.com/habibmrizki/week10/internal/service"
)

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Register(ctx *gin.Context) {
	var register dto.User
	if err := ctx.ShouldBindJSON(&register); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Msg:     "Format input tidak valid atau email salah",
		})
		return
	}

	if err := h.service.RegisterUser(register); err != nil {
		status := http.StatusBadRequest
		if err.Error() == "Email sudah terdaftar, gunakan email lain" {
			status = http.StatusConflict
		}
		ctx.JSON(status, dto.Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Data: gin.H{
			"email": register.Email,
		},
		Msg: "User berhasil didaftarkan",
	})
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var login dto.User
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Msg:     "Format input tidak valid",
		})
		return
	}

	storedUser, err := h.service.LoginUser(login)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, dto.Response{
			Success: false,
			Msg:     err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, dto.Response{
		Success: true,
		Data: gin.H{
			"email": storedUser.Email,
		},
		Msg: h.service.FormatLoginMessage(login.Email),
	})
}
