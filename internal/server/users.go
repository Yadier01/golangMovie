package server

import (
	"net/http"
	"time"

	"fmt"
	"github.com/Yadier01/golangMovie/db"
	"github.com/Yadier01/golangMovie/pkg/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (server *Server) CreateUser(c *gin.Context) {
	var user user

	if err := c.ShouldBindJSON(&user); err != nil {
		util.NewError(c, http.StatusBadRequest, "error with data")
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 15)
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, "internalServerError")
	}
	userParam := db.CreateUserParams{
		Name:     user.Name,
		Password: string(password),
		Email:    user.Email,
	}
	err = server.Query.CreateUser(c, userParam)
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, "could not create User")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"succes": "Created User Sucessfully"})
}

type LoginUser struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (server *Server) LoginUser(c *gin.Context) {
	var loginUser LoginUser

	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, "could not bind json")
		return
	}

	user, err := server.Query.GetUserByName(c, loginUser.Name)
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, "user does not exit")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, "name or password is wrong")
		return
	}

	token, err := util.SignUserJWT(user.ID, user.Name)
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, "could not make token")
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
