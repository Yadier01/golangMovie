// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"database/sql"
	"time"
)

type Movie struct {
	ID          int64
	Title       string
	Description string
	Genre       string
	Showtime    time.Time
	Seats       int32
	Poster      string
}

type Reservation struct {
	ID      int64
	Userid  int64
	Movieid int64
}

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
	Isadmin  sql.NullBool
}
