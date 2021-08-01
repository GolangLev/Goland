package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/GolangLev/Goland"
	"github.com/GolangLev/Goland/pkg/repository"
)

const solt = "vkodh73823gd2uq"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.GeneratePassword(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GeneratePassword(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(solt)))
}
