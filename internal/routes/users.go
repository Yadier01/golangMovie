package internal

import (
	"net/http"

	"github.com/Yadier01/golangMovie/db"
	"github.com/Yadier01/golangMovie/pkg/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) CreateUser(c *gin.Context) {
	var u user

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "error with data"})
		return
	}

	//check if user exists
	usr, _ := server.Query.GetUserByEmail(c, u.Email)
	if usr.Email == u.Email {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(u.Password), 15)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal sever error"})
		return
	}

	userParam := db.CreateUserParams{
		Name:     u.Name,
		Password: string(password),
		Email:    u.Email,
	}
	err = server.Query.CreateUser(c, userParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create User"})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not bind json"})
		return
	}

	user, err := server.Query.GetUserByName(c, loginUser.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user does not exist"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "name or password is wrong"})
		return
	}

	token, err := util.SignUserJWT(user.ID, user.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not make token"})
	}
	c.SetCookie("Auth", token, 65, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{"token": token})
}
