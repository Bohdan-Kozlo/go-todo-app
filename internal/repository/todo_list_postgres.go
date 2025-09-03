package repository

import (
	"github.com/bohdan-kozlo/todo-app/internal/models"
	"gorm.io/gorm"
)

type TodoListPostgres struct {
	db *gorm.DB
}

func NewTodoListPostgres(db *gorm.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (s *TodoListPostgres) Create(userId int, list models.TodoList) (int, error) {
	err := s.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&list).Error; err != nil {
			return err
		}

		link := models.UserList{UserID: uint(userId), TodoListID: list.ID}
		if err := tx.Create(&link).Error; err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return int(list.ID), nil
}
