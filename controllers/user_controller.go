package controllers

import (
	"encoding/json"
	"fmt"

	"go-postgres/models"
	UserService "go-postgres/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Create User POST request")
	rw.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	userid := UserService.CreateUser(&user)

	response := models.Response{
		Data:    userid,
		Message: "User Created Successfully",
	}

	json.NewEncoder(rw).Encode(response)
}

func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Get All Users GET request")
	rw.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	users, err := UserService.GetAllUsers()

	if err != nil {
		log.Fatal("Failed to get users")
	}
	response := models.Response{
		Data:    users,
		Message: "Ok",
	}

	json.NewEncoder(rw).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getUser function with user id to retrieve a single user
	user, err := UserService.GetUser(int64(id))
	var res models.Response
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res = models.Response{
			Data:    nil,
			Message: err.Error(),
		}
	} else {
		w.WriteHeader(http.StatusOK)
		res = models.Response{
			Data:    user,
			Message: "OK",
		}
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the getUser function with user id to retrieve a single user
	_, err = UserService.DeleteUser(int64(id))
	var res models.Response
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = models.Response{
			Data:    nil,
			Message: err.Error(),
		}
	} else {
		w.WriteHeader(http.StatusOK)
		res = models.Response{
			Data:    id,
			Message: "User Deleted",
		}
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the userid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// create an empty user of type models.User
	var user models.User

	// decode the json request to user
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	// call update user to update the user
	updatedRows := UserService.UpdateUser(int64(id), user)

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := models.Response{
		Data:    int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
