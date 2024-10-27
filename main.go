package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/Yadier01/golangMovie/db"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type server struct {
	db    *sql.DB
	query *db.Queries
}

func newServer(conn *sql.DB) *server {
	return &server{
		db:    conn,
		query: db.New(conn),
	}
}
func main() {
	//viper config
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading .env file: %s \n", err)
	}
	connStr := viper.GetString("connStr")

	conn, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", connStr, err)
		os.Exit(1)
	}
	defer conn.Close()
	r := gin.Default()

	srv := newServer(conn)

	r.POST("/movie", srv.checkUsr(), srv.createMovie)

	r.POST("/auth/signup", srv.createUser)
	r.GET("/movie/:name", srv.getMovies)
	r.POST("/reservation", srv.addReservation)
	r.Run()
}

// Middleware
func (server *server) checkUsr() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Middleware activated")

		userIDStr := c.GetHeader("userId")
		if userIDStr == "" {
			newError(c, http.StatusNotFound, "missing user ID")
			c.Abort()
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			newError(c, http.StatusBadRequest, "invalid user ID format")
			c.Abort()
			return
		}
		_, err = server.query.GetUserByID(c, int32(userID))
		if err != nil {
			newError(c, http.StatusBadRequest, "User is not admin or does not exist")
			c.Abort()
			return
		}

		c.Next()
	}
}

type reservation struct {
	UserId  int32 `json:"userId"`
	MovieId int32 `json:"movieId"`
}

func (server *server) addReservation(c *gin.Context) {

	var res reservation
	if err := c.ShouldBindJSON(&res); err != nil {
		newError(c, http.StatusNotFound, "movie not found")
		return
	}
	resParams := db.CreateReservationParams{
		Userid:  res.MovieId,
		Movieid: res.MovieId,
	}
	err := server.query.CreateReservation(c, resParams)
	if err != nil {
		newError(c, http.StatusInternalServerError, "could not create reservation")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "reservation Created"})
}

func (server *server) getMovies(c *gin.Context) {
	param := c.Param("name")

	movie, err := server.query.GetMovie(c, param)
	if err != nil {
		newError(c, http.StatusNotFound, "movie not found")
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

func (server *server) createMovie(c *gin.Context) {
	var movie Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		fmt.Println("----------")
		fmt.Println(movie, err)
		newError(c, http.StatusBadRequest, "bad json :(")
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
	err := server.query.CreateMovie(c, movieParams)
	if err != nil {

		newError(c, http.StatusInternalServerError, "could not create movie")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "movie created Sucessfully"})
}
func newError(c *gin.Context, status int, errorMessage string) {
	c.JSON(status, gin.H{"error": errorMessage})
}

type user struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (server *server) createUser(c *gin.Context) {
	var user user

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad json"})
		return
	}

	fmt.Println(user)
	userParam := db.CreateUserParams{
		Name:     user.Name,
		Password: user.Password,
		Email:    user.Email,
	}

	err := server.query.CreateUser(c, userParam)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"succes": "Created User Sucessfully"})

}
