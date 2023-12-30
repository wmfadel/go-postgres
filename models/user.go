package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID       int64  `json:"id"`
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
