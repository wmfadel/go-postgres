package router

import (
	"go-postgres/controllers"
	"go-postgres/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/user", controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/api/user", controllers.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/api/user/{id}", controllers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/api/user/{id}", controllers.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/api/login", controllers.Login).Methods(http.MethodPost)
	subrouter := router.PathPrefix("protected").Subrouter()
	subrouter.Use(middleware.AuthMiddleware)
	subrouter.HandleFunc("/", controllers.ProtectedController).Methods(http.MethodGet)

	router.PathPrefix("/api").Handler(subrouter)
	return router
}
