package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Yadier01/golangMovie/db"
	"github.com/Yadier01/golangMovie/pkg/util"
	"github.com/gin-gonic/gin"
)

type reservation struct {
	UserId  int32 `json:"userId"`
	MovieId int32 `json:"movieId"`
}

func (server *Server) addReservation(c *gin.Context) {
	var res reservation
	if err := c.ShouldBindJSON(&res); err != nil {
		util.NewError(c, http.StatusNotFound, "movie not found")
		return
	}
	resParams := db.CreateReservationParams{
		Userid:  res.MovieId,
		Movieid: res.MovieId,
	}
	err := server.Query.CreateReservation(c, resParams)
	if err != nil {
		util.NewError(c, http.StatusInternalServerError, "could not create reservation")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "reservation Created"})
}

func (server *Server) getMovies(c *gin.Context) {
	param := c.Param("name")

	movie, err := server.Query.GetMovie(c, param)
	if err != nil {
		util.NewError(c, http.StatusNotFound, "movie not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": movie})
}

type Movie struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Genre       string    `json:"genre"`
	Showtime    time.Time `json:"showtime"`
	Seats       int32     `json:"seats"`
	Poster      string    `json:"poster"`
}

func (server *Server) CreateMovie(c *gin.Context) {
	var movie Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		fmt.Println("----------")
		fmt.Println(movie, err)
		util.NewError(c, http.StatusBadRequest, "bad json :(")
		return
	}

	movieParams := db.CreateMovieParams{
		Title:       movie.Title,
		Description: movie.Description,
		Genre:       movie.Genre,
		Showtime:    movie.Showtime,
		Seats:       movie.Seats,
		Poster:      movie.Poster,
	}
	err := server.Query.CreateMovie(c, movieParams)
	if err != nil {

		util.NewError(c, http.StatusInternalServerError, "could not create movie")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "movie created Sucessfully"})
}
