package main

import (
	"fmt"
	"go-postgres/database"
	"go-postgres/router"
	"go-postgres/utils"
	"log"
	"net/http"
)

func main() {

	connectionString, err := utils.GetConnectionString()
	if err != nil {
		log.Fatal("Failed to get connection string from environment", err)
	}
	database.Connect(connectionString)
	defer database.Instance.Close()
	r := router.Router()
	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
