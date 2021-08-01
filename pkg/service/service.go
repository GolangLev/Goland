package service

import (
	todo "github.com/GolangLev/Goland"
	"github.com/GolangLev/Goland/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService /*Внедрение зависимостей для общения с репозиторием*/
func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
