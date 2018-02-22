package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"ReviewService/models"
)

var tokenKey = "myKey"

func CreateToken(user *models.User) string {
	claims := make(jwt.MapClaims)
	claims["userId"]=user.Id
	claims["exp"]=time.Now().Add(time.Hour*480).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(tokenKey))
	return tokenString
}