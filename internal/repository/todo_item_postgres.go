package repository

import (
	"github.com/bohdan-kozlo/todo-app/internal/models"
	"gorm.io/gorm"
)

type TodoItemPostgres struct {
	db *gorm.DB
}

func NewTodoItemPostgres(db *gorm.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(listId int, item models.TodoItem) (int, error) {
	item.ListID = uint(listId)
	if err := r.db.Create(&item).Error; err != nil {
		return 0, err
	}
	return int(item.ID), nil
}

func (r *TodoItemPostgres) GetAll(listId int) ([]models.TodoItem, error) {
	var items []models.TodoItem
	if err := r.db.Where("list_id = ?", listId).Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TodoItemPostgres) GetById(listId, itemId int) (models.TodoItem, error) {
	var item models.TodoItem
	if err := r.db.Where("list_id = ? AND id = ?", listId, itemId).First(&item).Error; err != nil {
		return models.TodoItem{}, err
	}
	return item, nil
}

func (r *TodoItemPostgres) Update(updates map[string]interface{}, item *models.TodoItem) error {
	return r.db.Model(&item).Updates(updates).Error
}

func (r *TodoItemPostgres) Delete(item *models.TodoItem) error {
	return r.db.Delete(item).Error
}
