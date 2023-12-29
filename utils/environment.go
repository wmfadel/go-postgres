package utils

import (
	"os"

	"github.com/joho/godotenv"
)

func GetConnectionString() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	return os.Getenv("POSTGRES_URL"), nil
}
