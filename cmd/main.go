package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Msg     string `json:"msg"`
}

var users = make(map[string]User)

func validateAccount(account User) error {
	if len(account.Email) == 0 {
		return errors.New("email tidak boleh kosong")
	}
	if len(account.Password) < 8 {
		return errors.New("password tidak boleh kurang dari 8 karakter")
	}
	return nil
}

func main() {
	r := gin.Default()

	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", func(ctx *gin.Context) {
			var register User
			if err := ctx.ShouldBindJSON(&register); err != nil {
				ctx.JSON(http.StatusBadRequest, Response{
					Success: false,
					Msg:     "Format input tidak valid atau email salah",
				})
				return
			}

			if err := validateAccount(register); err != nil {
				ctx.JSON(http.StatusBadRequest, Response{
					Success: false,
					Msg:     err.Error(),
				})
				return
			}

			if _, exists := users[register.Email]; exists {
				ctx.JSON(http.StatusConflict, Response{
					Success: false,
					Msg:     "Email sudah terdaftar, gunakan email lain",
				})
				return
			}

			users[register.Email] = register

			ctx.JSON(http.StatusCreated, Response{
				Success: true,
				Data: gin.H{
					"email": register.Email,
				},
				Msg: "User berhasil didaftarkan",
			})
		})

		authGroup.POST("/login", func(ctx *gin.Context) {
			var login User
			if err := ctx.ShouldBindJSON(&login); err != nil {
				ctx.JSON(http.StatusBadRequest, Response{
					Success: false,
					Msg:     "Format input tidak valid",
				})
				return
			}

			storedUser, exists := users[login.Email]
			if !exists || storedUser.Password != login.Password {
				ctx.JSON(http.StatusUnauthorized, Response{
					Success: false,
					Msg:     "Email atau password salah",
				})
				return
			}

			ctx.JSON(http.StatusOK, Response{
				Success: true,
				Data: gin.H{
					"email": storedUser.Email,
				},
				Msg: fmt.Sprintf("Anda berhasil login dengan email %s", login.Email),
			})
		})
	}

	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, Response{
			Msg:     "rute salah",
			Success: false,
		})
	})

	r.Run(":3000")
}
