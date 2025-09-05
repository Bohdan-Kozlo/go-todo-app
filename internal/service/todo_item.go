package service

import (
	"github.com/bohdan-kozlo/todo-app/internal/models"
	"github.com/bohdan-kozlo/todo-app/internal/repository"
)

type TodoItemService struct {
	itemRepo repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(itemRepo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{itemRepo: itemRepo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId, listId int, item models.TodoItem) (int, error) {
	if _, err := s.listRepo.GetById(userId, listId); err != nil {
		return 0, err
	}
	return s.itemRepo.Create(listId, item)
}

func (s *TodoItemService) GetAll(userId, listId int) ([]models.TodoItem, error) {
	if _, err := s.listRepo.GetById(userId, listId); err != nil {
		return nil, err
	}
	return s.itemRepo.GetAll(listId)
}

func (s *TodoItemService) GetById(userId, listId, itemId int) (models.TodoItem, error) {
	if _, err := s.listRepo.GetById(userId, listId); err != nil {
		return models.TodoItem{}, err
	}
	return s.itemRepo.GetById(listId, itemId)
}

func (s *TodoItemService) Update(userId, listId, itemId int, input models.UpdateItemInput) error {
	if _, err := s.listRepo.GetById(userId, listId); err != nil {
		return err
	}
	item, err := s.itemRepo.GetById(listId, itemId)
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
	if input.Completed != nil {
		updates["completed"] = *input.Completed
	}
	if len(updates) == 0 {
		return nil
	}

	return s.itemRepo.Update(updates, &item)
}

func (s *TodoItemService) Delete(userId, listId, itemId int) error {
	if _, err := s.listRepo.GetById(userId, listId); err != nil {
		return err
	}
	item, err := s.itemRepo.GetById(listId, itemId)
	if err != nil {
		return err
	}
	return s.itemRepo.Delete(&item)
}
