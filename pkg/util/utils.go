package util

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// change this obv, should be in .env file also
var mySigningKey = []byte("AllYourBase")

func SignUserJWT(id int64, name string) (string, error) {
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

type claims struct {
	Id   int64
	Name string
	Exp  int64
}

func CheckUserJWT(tokenString string) (*claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return mySigningKey, nil
	})
	if err != nil {
		return nil, errors.New("could not parse token")
	}
	if token.Valid {
		return &claims{
			Id:   int64(token.Claims.(jwt.MapClaims)["id"].(float64)),
			Name: token.Claims.(jwt.MapClaims)["name"].(string),
			Exp:  int64(token.Claims.(jwt.MapClaims)["exp"].(float64)),
		}, nil
	}
	return nil, err
}
