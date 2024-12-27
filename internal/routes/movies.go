package internal

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Yadier01/golangMovie/db"
	"github.com/gin-gonic/gin"
)

func (server *Server) getMovies(c *gin.Context) {
	movies, err := server.Query.GetMovies(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch movies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}
func (server *Server) getMovie(c *gin.Context) {
	param := c.Param("name")

	movie, err := server.Query.GetMovie(c, param)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "this movie does not exits"})
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
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad json"})
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
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "could not create movie"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "movie created Sucessfully"})
}
