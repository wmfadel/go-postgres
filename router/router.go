package router

import (
	"go-postgres/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/user", middleware.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/user", middleware.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/user/{id}", middleware.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/api/user/{id}", middleware.DeleteUser).Methods(http.MethodDelete)
	return router
}
