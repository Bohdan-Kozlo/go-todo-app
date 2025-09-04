package service

import (
	"github.com/bohdan-kozlo/todo-app/internal/models"
	"github.com/bohdan-kozlo/todo-app/internal/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list models.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]models.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (models.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Delete(userId, listId int) error {
	list, err := s.repo.GetById(userId, listId)
	if err != nil {
		return err
	}

	return s.repo.Delete(&list)
}

func (s *TodoListService) Update(userId, listId int, input models.UpdateListInput) error {
	list, err := s.GetById(userId, listId)
	if err != nil {
		return err
	}

	updates := map[string]interface{}{}
	if input.Title != nil {
		updates["title"] = *input.Title
	}
	if input.Description != nil {
		updates["description"] = *input.Description
	}
	if len(updates) == 0 {
		return nil
	}

	return s.repo.Update(updates, &list)
}
