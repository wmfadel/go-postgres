package controllers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"go-postgres/database"
	"go-postgres/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// response format
type response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type RequestError struct {
	StatusCode int

	Err error
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

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

	userid := createUser(&user)

	response := response{
		Data:    userid,
		Message: "User Created Successfully",
	}

	json.NewEncoder(rw).Encode(response)
}

func GetAllUsers(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("Handling Get All Users GET request")
	rw.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	users, err := getAllUsers()

	if err != nil {
		log.Fatal("Failed to get users")
	}
	response := response{
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
	user, err := getUser(int64(id))
	var res response
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		res = response{
			Data:    nil,
			Message: err.Error(),
		}
	} else {
		w.WriteHeader(http.StatusOK)
		res = response{
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
	_, err = deleteUser(int64(id))
	var res response
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		res = response{
			Data:    nil,
			Message: err.Error(),
		}
	} else {
		w.WriteHeader(http.StatusOK)
		res = response{
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
	updatedRows := updateUser(int64(id), user)

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := response{
		Data:    int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func createUser(user *models.User) int64 {
	user.HashPassword()
	sqlStatement := `INSERT INTO users (name, password, location, age) VALUES ($1, $2, $3, $4) RETURNING userid`
	var userid int64
	err := database.Instance.QueryRow(sqlStatement, user.Name, user.Password, user.Location, user.Age).Scan(&userid)
	if err != nil {
		log.Fatalf("Failed to create user.  %v", err)
	}
	fmt.Println("Created new user")
	return userid
}

func updateUser(id int64, user models.User) int64 {
	user.HashPassword()
	sqlStatement := `UPDATE users SET name=$2, password=$3, location=$4, age=$5 WHERE userid=$1`
	res, err := database.Instance.Exec(sqlStatement, id, user.Name, user.Password, user.Location, user.Age)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}

func getAllUsers() ([]models.User, error) {

	rows, err := database.Instance.Query(`SELECT userid, name, age, location FROM users`)
	if err != nil {
		log.Fatal("Failed to query users")
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		users = append(users, user)
	}

	return users, err
}

func getUser(Id int64) (models.User, error) {

	row := database.Instance.QueryRow(`SELECT userid, name, age, location FROM users WHERE userid = $1`, Id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return user, errors.New("no rows found")
	case nil:
		return user, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty user on error
	return user, err
}

func deleteUser(Id int64) (int64, error) {

	res, err := database.Instance.Exec(`DELETE FROM users WHERE userid=$1`, Id)

	if err != nil {
		return 0, err
	}

	rowsCount, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsCount != 1 {
		return 0, errors.New("failed to delete user")
	}

	return Id, nil
}
