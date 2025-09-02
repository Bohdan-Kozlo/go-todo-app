package database

import (
	"os"

	"github.com/bohdan-kozlo/todo-app/internal/models"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		logrus.Fatal("database url is not provided")
	}

	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		logrus.Fatal("failed to connect to database: ", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.TodoList{},
		&models.TodoItem{},
		&models.UserList{},
	)
	if err != nil {
		logrus.Fatal("failed to migrate database: ", err)
	}

	return db
}
