package service

import (
	"errors"
	"fmt"

	"github.com/habibmrizki/week10/internal/dto"
)

type AuthService struct {
	users map[string]dto.User
}

func NewAuthService() *AuthService {
	return &AuthService{
		users: make(map[string]dto.User),
	}
}

func (s *AuthService) ValidateAccount(account dto.User) error {
	if len(account.Email) == 0 {
		return errors.New("email tidak boleh kosong")
	}
	if len(account.Password) < 8 {
		return errors.New("password tidak boleh kurang dari 8 karakter")
	}
	return nil
}

func (s *AuthService) RegisterUser(account dto.User) error {
	if err := s.ValidateAccount(account); err != nil {
		return err
	}

	if _, exists := s.users[account.Email]; exists {
		return errors.New("Email sudah terdaftar, gunakan email lain")
	}

	s.users[account.Email] = account
	return nil
}

func (s *AuthService) LoginUser(account dto.User) (dto.User, error) {
	storedUser, exists := s.users[account.Email]
	if !exists || storedUser.Password != account.Password {
		return dto.User{}, errors.New("Email atau password salah")
	}

	return storedUser, nil
}

func (s *AuthService) FormatLoginMessage(email string) string {
	return fmt.Sprintf("Anda berhasil login dengan email %s", email)
}
