package internal

import (
	"database/sql"

	"github.com/Yadier01/golangMovie/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	Db    *sql.DB
	Query *db.Queries
}

func NewServer(conn *sql.DB) *Server {
	return &Server{
		Db:    conn,
		Query: db.New(conn),
	}
}

func (srv *Server) New() {
	r := gin.Default()
	r.POST("/movies", srv.CheckUsr(), srv.IsAdmin(), srv.CreateMovie)

	r.POST("/auth/signup", srv.CreateUser)
	r.POST("auth/login", srv.LoginUser)

	r.GET("/movies/:name", srv.getMovie)
	r.GET("/movies", srv.getMovies)

	r.POST("movies/reservations/:name", srv.CheckUsr(), srv.addReservation)
	r.GET("movies/reservations/", srv.CheckUsr(), srv.getReservations)
	r.DELETE("movies/reservations/:resId", srv.CheckUsr(), srv.removeReservation)
	r.Run()
}
