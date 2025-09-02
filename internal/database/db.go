package database

import (
	"log"
	"os"

	"github.com/bohdan-kozlo/todo-app/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("database url is not provided")
	}

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.TodoList{},
		&models.TodoItem{},
		&models.UserList{},
	)
	if err != nil {
		log.Fatal("failed to migrate database: ", err)
	}

	return db
}
