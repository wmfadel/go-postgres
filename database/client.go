package database

import (
	"errors"
	"fmt"
	"go-postgres/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Instance *gorm.DB

func Connect(connectionString string) {

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.AutoMigrate(models.User{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connection established")
	Instance = db
}

func Disconnect() error {
	if Instance == nil {
		return errors.New("connection already closed")
	}
	db, err := Instance.DB()

	if err != nil {
		return err
	}

	db.Close()
	return nil
}
