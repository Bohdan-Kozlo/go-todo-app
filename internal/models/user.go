package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Lists []TodoList `gorm:"many2many:user_lists;joinForeignKey:UserID;joinReferences:TodoListID"`
}

type UserList struct {
	UserID     uint `gorm:"primaryKey"`
	TodoListID uint `gorm:"primaryKey"`
}
