MIGRATE=goose
DRIVER=postgres
DIR="./db/migrations/"
.PHONY: up down 

ifneq (,$(wildcard .env))
    include .env
endif

up:
	$(MIGRATE) $(DRIVER) $(DB_URL) -dir ${DIR}  up 

down:
	$(MIGRATE) $(DRIVER) $(DB_URL) -dir ${DIR} down 

# Create a new migration file
create:
	$(MIGRATE) create $(DB_URL) migration_name_here
