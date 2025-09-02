package repository

import (
	"github.com/bohdan-kozlo/todo-app/internal/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
