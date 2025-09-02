package repository

import (
	"github.com/bohdan-kozlo/todo-app/internal/models"
	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(user.ID), nil
}
