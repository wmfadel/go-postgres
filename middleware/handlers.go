package middleware

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"go-postgres/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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

func createConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to open .env", err)
	}

	connectionString := os.Getenv("POSTGRES_URL")

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection established")
	return db
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

func createUser(user *models.User) int64 {
	db := createConnection()

	defer db.Close()

	sqlStatement := `INSERT INTO users (name, location, age) VALUES ($1, $2, $3) RETURNING userid`
	var userid int64
	err := db.QueryRow(sqlStatement, user.Name, user.Location, user.Age).Scan(&userid)
	if err != nil {
		log.Fatalf("Failed to create user.  %v", err)
	}
	fmt.Println("Created new user")
	return userid
}

func getAllUsers() ([]models.User, error) {
	db := createConnection()
	defer db.Close()
	rows, err := db.Query(`SELECT * FROM users`)
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
	db := createConnection()
	defer db.Close()
	row := db.QueryRow(`SELECT * FROM users WHERE userid = $1`, Id)

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
