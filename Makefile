.PHONY: createdb dropdb
postgres:
	sudo  docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:17.2-alpine
createdb:
	# Create a database in the Docker container
	sudo docker exec -it postgres17 createdb --username=root --owner=root movie_golang

dropdb:
	# Drop the database in the Docker container
	sudo docker exec -it postgres17 dropdb movie_golang
