-- name: GetUserByID :one
SELECT * FROM users
WHERE users.id = $1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE users.name = $1;

-- name: CreateUser :exec
INSERT INTO users (
  name, email, password
) VALUES ( $1, $2, $3 );

-- name: GetUserByEmail :one
SELECT *
  FROM users 
  WHERE email = $1;


-- name: CreateMovie :exec
INSERT INTO movies (
 title, description, genre, showtime, seats, poster
) VALUES ( $1, $2, $3,$4,$5,$6);

-- name: GetMovie :one
SELECT * FROM movies
WHERE movies.title = $1;

-- name: GetMovieById :one
SELECT *
  FROM movies
  WHERE id = $1;

-- name: GetMovies :many
SELECT * FROM movies
LIMIT 10;






-- name: CreateReservation :exec
INSERT INTO reservations (
  userid, movieid
) VALUES ( $1, $2);

-- name: GetReservations :many
SELECT * FROM reservations as r
  INNER JOIN movies
  ON r.movieid = movies.id
  WHERE r.userid = $1;

-- name: GetReservation :one
SELECT *
  FROM reservations as r
  WHERE r.id = $1;
-- name: RemoveReservation :exec
DELETE FROM reservations as r
  WHERE r.id= $1 AND r.userid = $2;
 






