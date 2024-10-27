package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Yadier01/golangMovie/pkg/util"
	"github.com/gin-gonic/gin"
)

// Middleware
func (server *Server) CheckUsr() gin.HandlerFunc {

	return func(c *gin.Context) {

		userIDStr := c.GetHeader("userId")
		if userIDStr == "" {
			util.NewError(c, http.StatusNotFound, "missing user ID")
			c.Abort()
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			util.NewError(c, http.StatusBadRequest, "invalid user ID format")
			c.Abort()
			return
		}
		user, err := server.Query.GetUserByID(c, int32(userID))
		if err != nil {
			util.NewError(c, http.StatusBadRequest, "User is not admin or does not exist")
			c.Abort()
			return
		}
		c.Set("id", user.ID)
		c.Next()
	}
}

func (server *Server) IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, exist := c.Get("id")
		if !exist {
			util.NewError(c, http.StatusNotFound, "missing user ID")
			c.Abort()
			return
		}

		// Retrieve user by ID
		usr, err := server.Query.GetUserByID(c, id.(int32))
		if err != nil {
			util.NewError(c, http.StatusInternalServerError, "could not fetch user")
			c.Abort()
			return
		}

		// Check if the user is an admin
		if usr.Isadmin.Valid && usr.Isadmin.Bool {
			fmt.Println("user IS admin")
		} else {
			util.NewError(c, http.StatusForbidden, "user is not an admin")
			c.Abort()
			return
		}
	}
}
