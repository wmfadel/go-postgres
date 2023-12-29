package router

import (
	"go-postgres/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/user", controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/user", controllers.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/user/{id}", controllers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods(http.MethodDelete)
	return router
}
