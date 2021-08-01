package service

import (
	todo "github.com/GolangLev/Goland"
	"github.com/GolangLev/Goland/pkg/repository"
)

type TodoServiceItem struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoServiceItem(repo repository.TodoItem, listRepo repository.TodoList) *TodoServiceItem {
	return &TodoServiceItem{repo: repo, listRepo: listRepo}
}

func (s *TodoServiceItem) Create(userId int, listId int, item todo.TodoItem) (int, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		//Если такого списка или он не пренадлежит пользователю, то возращаем ошибку
		return 0, err
	}

	return s.repo.Create(listId, item)
}

func (s *TodoServiceItem) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(userId, listId)
}

func (s *TodoServiceItem) GetById(userId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(userId, itemId)
}

func (s *TodoServiceItem) Delete(userId, itemId int) error {
	return s.repo.Delete(userId, itemId)
}

func (s *TodoServiceItem) Update(userId, itemId int, input todo.UpdateItemInput) error {
	return s.repo.Update(userId, itemId, input)
}
