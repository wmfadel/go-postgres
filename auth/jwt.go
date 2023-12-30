package auth

import (
	"errors"
	"fmt"
	"go-postgres/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	UserId int64  `json:"userid"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(userid int64, name string, email string) (tokenString string, expiresAt int64, err error) {
	tokenDuration, err := utils.GetJWTExpiaryTime()
	if err != nil {
		tokenDuration = 1
		fmt.Println("Failed to get jwt duration using default value 1")
	}
	expirationTime := time.Now().Add(time.Duration(tokenDuration) * time.Hour)
	expiresAt = expirationTime.Unix()
	claims := &JWTClaim{
		UserId: userid,
		Name:   name,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateToken(signedToken string) (err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return

}
