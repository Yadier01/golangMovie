package server

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
	r.POST("/movie", srv.CheckUsr(), srv.IsAdmin(), srv.CreateMovie)

	r.POST("/auth/signup", srv.CreateUser)
	r.GET("/movie/:name", srv.getMovies)
	r.POST("/reservation", srv.addReservation)
	r.Run()
}
