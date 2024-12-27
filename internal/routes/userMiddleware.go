package internal

import (
	"net/http"

	"github.com/Yadier01/golangMovie/pkg/util"
	"github.com/gin-gonic/gin"
)

// Middleware
func (server *Server) CheckUsr() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Auth")

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		claims, err := util.CheckUserJWT(token)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "token is not valid"})
			c.Abort()
			return
		}

		_, err = server.Query.GetUserByID(c, claims.Id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "user does not exits"})
			c.Abort()
			return
		}
		c.Set("id", claims.Id)
		c.Set("name", claims.Name)
		c.Next()
	}
}

func (server *Server) IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("Auth")
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "missing token"})
			c.Abort()
			return
		}

		claims, err := util.CheckUserJWT(token)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"message": "invalid token"})
			c.Abort()
			return
		}

		usr, err := server.Query.GetUserByID(c, claims.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch user"})
			c.Abort()
			return
		}

		// Check if the user is an admin
		if usr.Isadmin.Valid && usr.Isadmin.Bool {
			c.Next()
			return
		} else {
			c.JSON(http.StatusForbidden, gin.H{"message": "User is not admin"})
			c.Abort()
			return
		}
	}
}
