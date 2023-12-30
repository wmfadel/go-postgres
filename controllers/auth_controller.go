package controllers

import (
	"encoding/json"
	"fmt"
	"go-postgres/auth"
	"go-postgres/models"
	UserService "go-postgres/services"
	"log"
	"net/http"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var loginRequest models.LoginRequest

	err := json.NewDecoder(r.Body).Decode(&loginRequest)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getUser function with user id to retrieve a single user
	user, err := UserService.GetUserByEmail(loginRequest.Email)
	var res models.Response
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		res = models.Response{
			Data:    nil,
			Message: err.Error(),
		}
		json.NewEncoder(rw).Encode(res)
		return
	}
	fmt.Println("user to be checked", user.Password)
	err = user.CheckPassword(loginRequest.Password)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		res = models.Response{
			Data:    nil,
			Message: err.Error(),
		}
		json.NewEncoder(rw).Encode(res)
		return
	}

	token, expiary, err := auth.GenerateJWT(user.ID, user.Name, user.Email)

	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(rw).Encode(models.RequestError{
			StatusCode: http.StatusBadRequest,
			Err:        "Login Failed, Try again later",
		})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(models.AuthenticationResponse{
		User:      user,
		Token:     token,
		ExpiresAt: expiary,
		Message:   "ok",
	})
}
