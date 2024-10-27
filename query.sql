-- name: GetUserByID :one
SELECT * FROM users 
WHERE users.id = $1;

-- name: CreateUser :exec 
INSERT INTO users (
  name, email, password 
) VALUES ( $1, $2, $3 );

-- name: CreateMovie :exec
INSERT INTO movies (
 title, description, genre, showtime, seats, poster 
) VALUES ( $1, $2, $3,$4,$5,$6);

-- name: GetMovie :one
SELECT * 
  FROM movies 
  WHERE movies.title = $1;

-- name: CreateReservation :exec
INSERT INTO reservations (
  userid, movieid
) VALUES ( $1, $2);

