package controllers

import (
	"encoding/json"
	"go-postgres/models"
	"net/http"
)

func ProtectedController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(models.Response{
		Data:    "ok",
		Message: "ok",
	})
}
