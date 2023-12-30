package UserService

import (
	"database/sql"
	"errors"
	"fmt"
	"go-postgres/database"
	"go-postgres/models"
	"log"
)

func CreateUser(user *models.User) int64 {
	user.HashPassword()
	sqlStatement := `INSERT INTO users (name, password, location, age, email) VALUES ($1, $2, $3, $4, $5) RETURNING userid`
	var userid int64
	err := database.Instance.QueryRow(sqlStatement, user.Name, user.Password, user.Location, user.Age, user.Email).Scan(&userid)
	if err != nil {
		log.Fatalf("Failed to create user.  %v", err)
	}
	fmt.Println("Created new user")
	return userid
}

func UpdateUser(id int64, user models.User) int64 {
	user.HashPassword()
	sqlStatement := `UPDATE users SET name=$2, password=$3, location=$4, age=$5, email-$6 WHERE userid=$1`
	res, err := database.Instance.Exec(sqlStatement, id, user.Name, user.Password, user.Location, user.Age, user.Email)

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

func GetAllUsers() ([]models.User, error) {

	rows, err := database.Instance.Query(`SELECT userid, name, age, location, email FROM users`)
	if err != nil {
		log.Fatal("Failed to query users")
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Age, &user.Location, &user.Email)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		users = append(users, user)
	}

	return users, err
}

func GetUser(Id int64) (models.User, error) {

	row := database.Instance.QueryRow(`SELECT userid, name, age,location, email  FROM users WHERE userid = $1`, Id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location, &user.Email)

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

func GetUserByEmail(email string) (models.User, error) {

	row := database.Instance.QueryRow(`SELECT * FROM users WHERE email = $1`, email)

	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Location, &user.Password, &user.Email)

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

func DeleteUser(Id int64) (int64, error) {

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
