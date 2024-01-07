package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	//! TODO: check here the ID may fuckup everything here
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Location string `json:"location"`
	Age      int64  `json:"age"`
}

func (user *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate: Hashing password before creating new user")
	user.HashPassword()
	fmt.Println("BeforeCreate: Password hashing completed")
	return
}

func (user *User) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterCreate: New user created with ID", user.ID)
	return
}
