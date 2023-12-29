package database

import (
	"database/sql"
	"fmt"
)

var Instance *sql.DB

func Connect(connectionString string) {

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connection established")
	Instance = db
}
