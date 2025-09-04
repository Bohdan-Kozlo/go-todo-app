package models

import "time"

type TodoList struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	Users []User `gorm:"many2many:user_lists;joinForeignKey:TodoListID;joinReferences:UserID"`

	Items []TodoItem `gorm:"foreignKey:ListID;constraint:OnDelete:CASCADE;"`
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}
