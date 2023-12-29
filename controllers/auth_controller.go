package controllers

import (
	"net/http"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	// fmt.Println("Handling Create User POST request")
	// rw.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	// rw.Header().Set("Access-Control-Allow-Origin", "*")
	// rw.Header().Set("Access-Control-Allow-Methods", "POST")
	// rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// var user models.User

	// err := json.NewDecoder(r.Body).Decode(&user)

	// if err != nil {
	// 	log.Fatalf("Unable to decode the request body.  %v", err)
	// }

	// userid := createUser(&user)

	// response := models.Response{
	// 	Data:    userid,
	// 	Message: "User Created Successfully",
	// }

	// json.NewEncoder(rw).Encode(response)
}
