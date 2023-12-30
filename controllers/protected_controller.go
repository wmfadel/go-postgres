package controllers

import (
	"encoding/json"
	"go-postgres/models"
	"net/http"
)

func ProtectedController(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.Response{
		Data:    "ok",
		Message: "ok",
	})
}
