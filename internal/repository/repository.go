package repository

import (
	"github.com/bohdan-kozlo/todo-app/internal/models"
	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username string) (models.User, error)
}

type TodoList interface {
	Create(userId int, list models.TodoList) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Delete(list *models.TodoList) error
	Update(updates map[string]interface{}, list *models.TodoList) error
}

type TodoItem interface {
	Create(listId int, item models.TodoItem) (int, error)
	GetAll(listId int) ([]models.TodoItem, error)
	GetById(listId, itemId int) (models.TodoItem, error)
	Update(updates map[string]interface{}, item *models.TodoItem) error
	Delete(item *models.TodoItem) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
