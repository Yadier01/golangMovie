-- +goose Up
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(40) NOT NULL,
    isadmin BOOLEAN 
);


CREATE TABLE movies (
 id SERIAL PRIMARY KEY, 
 title VARCHAR(255) NOT NULL,
 description VARCHAR(255) NOT NULL,
 genre VARCHAR(20) NOT NULL,
 showtime TIMESTAMP NOT NULL,
 seats INTEGER NOT NULL,
 poster VARCHAR(255) NOT NULL
);

CREATE TABLE reservations (
  id SERIAL PRIMARY KEY,   
  userid INT NOT NULL,
  movieid INT NOT NULL,
  CONSTRAINT fk_users FOREIGN KEY(userid) REFERENCES users(id),
  CONSTRAINT fk_movies FOREIGN KEY(movieid) REFERENCES movies(id)
);
-- +goose Down
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS reservations;
