package internal

import (
	"net/http"
	"strconv"

	"github.com/Yadier01/golangMovie/db"
	"github.com/gin-gonic/gin"
)

func (s *Server) addReservation(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Please log in"})
		return
	}
	movieName := c.Param("name")
	movie, err := s.Query.GetMovie(c, movieName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get movie"})
		return
	}

	resParams := db.CreateReservationParams{
		Userid:  id.(int64),
		Movieid: movie.ID,
	}

	err = s.Query.CreateReservation(c, resParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create reservation"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "reservation Created"})
}
func (s *Server) getReservations(c *gin.Context) {
	userId, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Please log in or register"})
		return
	}
	reservations, err := s.Query.GetReservations(c, userId.(int64))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not get reservations"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
func (s *Server) removeReservation(c *gin.Context) {
	resId, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": "Please log in or register"})
		return
	}

	movieId, err := strconv.ParseInt(c.Param("resId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "invalid id"})
		return
	}
	// check if reservation exits
	res, err := s.Query.GetReservation(c, movieId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "That Reservation does not exists"})
		return
	}

	if res.Userid != resId.(int64) {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	params := db.RemoveReservationParams{
		ID:     movieId,
		Userid: resId.(int64),
	}
	err = s.Query.RemoveReservation(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not delete"})
		return
	}

	c.JSON(http.StatusMovedPermanently, gin.H{"message": "deleted"})
}
