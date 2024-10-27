package util

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func NewError(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, gin.H{"error": errorMessage})
}
func SignUserJWT(id int32, name string) (jwtToken string, err error) {
	//change this obv
	var mySigningKey = []byte("AllYourBase")

	// Create the Claims
	claims := jwt.MapClaims{
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		"id":   id,
		"name": name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return ss, nil
}

// fix this
func CheckUserJWT(tokenString string) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})

}
