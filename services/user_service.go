package UserService

import (
	"errors"
	"fmt"
	"go-postgres/database"
	"go-postgres/models"
	"log"

	"gorm.io/gorm"
)

func CreateUser(user *models.User) uint {
	tx := database.Instance.Create(user)
	if tx.Error != nil {
		log.Fatalf("Failed to create user.  %v", tx.Error)
	}
	fmt.Println("Created new user")
	return user.ID
}

func UpdateUser(id int64, user models.User) models.User {
	var updatedUser models.User
	tx := database.Instance.Model(&updatedUser).Where("id = ?", id).Updates(&user)
	if tx.Error != nil {
		log.Fatalf("Unable to execute the query. %v", tx.Error)
	}

	if tx.RowsAffected == 0 {
		log.Fatalf("Error while checking the affected rows 0")
	}

	return updatedUser
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	tx := database.Instance.Find(&users)
	if tx.Error != nil {
		log.Fatal("Failed to query users")
		return nil, tx.Error
	}

	return users, tx.Error
}

func GetUser(Id uint) (models.User, error) {
	var user models.User
	tx := database.Instance.First(&user, Id)

	if tx.Error != nil {
		return user, tx.Error
	}
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("user not found")
	}

	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {

	var user models.User
	tx := database.Instance.Where("email = ?", email).First(&user)

	if tx.Error != nil {
		return user, tx.Error
	}
	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("user not found")
	}

	return user, nil
}

func DeleteUser(Id int64) (models.User, error) {
	var user models.User
	tx := database.Instance.Delete(&user, Id)

	if tx.Error != nil {
		return user, tx.Error
	}

	if tx.RowsAffected != 1 {
		return user, errors.New("didn't find the user to delete")
	}

	return user, nil
}
