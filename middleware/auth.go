package middleware

import (
	"encoding/json"
	"fmt"
	"go-postgres/auth"
	"go-postgres/models"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		fmt.Println("checking token", token)
		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.RequestError{
				StatusCode: http.StatusUnauthorized,
				Err:        "missing token",
			})
			return
		}

		err := auth.ValidateToken(strings.Split(token, " ")[1])
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(models.RequestError{
				StatusCode: http.StatusUnauthorized,
				Err:        "invalid token",
			})
			return
		}

		next.ServeHTTP(w, r)
	})
}
