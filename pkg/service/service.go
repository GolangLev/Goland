package service

import "github.com/GolangLev/Goland/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
