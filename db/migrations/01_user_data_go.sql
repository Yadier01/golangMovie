-- +goose Up
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    isadmin BOOLEAN
);

CREATE TABLE movies (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    genre TEXT NOT NULL,
    showtime TIMESTAMP NOT NULL,
    seats INTEGER NOT NULL,
    poster TEXT NOT NULL
);

CREATE TABLE reservations (
    id BIGSERIAL PRIMARY KEY,
    userid BIGINT NOT NULL,
    movieid BIGINT NOT NULL,
    CONSTRAINT fk_users FOREIGN KEY (userid) REFERENCES users (id),
    CONSTRAINT fk_movies FOREIGN KEY (movieid) REFERENCES movies (id)
);

-- +goose Down
DROP TABLE IF EXISTS reservations;

DROP TABLE IF EXISTS movies;

DROP TABLE IF EXISTS users;
