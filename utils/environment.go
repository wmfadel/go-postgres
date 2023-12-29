package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetConnectionString() (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	return os.Getenv("POSTGRES_URL"), nil
}

func GetJWTExpiaryTime() (int, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return 0, err
	}
	expiray, err := strconv.Atoi(os.Getenv("JWT_EXPIARY"))
	if err != nil {
		return 0, err
	}
	return expiray, nil

}
