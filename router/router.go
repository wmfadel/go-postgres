package router

import (
	"go-postgres/controllers"
	"go-postgres/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/admin/user", controllers.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/admin/user", controllers.GetAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/admin/user/{id}", controllers.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/admin/user/{id}", controllers.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/admin/user/{id}", controllers.DeleteUser).Methods(http.MethodDelete)
	router.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	subrouter := router.PathPrefix("/api").Subrouter()
	subrouter.Use(middleware.AuthMiddleware)
	subrouter.HandleFunc("/protected", controllers.ProtectedController).Methods(http.MethodGet)

	return router
}
