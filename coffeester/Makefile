

include .env


stop_containers: 
	@echo "Stopping other docker container"
	if [ $$(docker ps -q) ]; then \
		echo "found and stopped containers"; \
		docker stop $$(docker s -q); \
	else \
		echo "no containers running..."; \
	fi 

create_container:
	docker run --name ${DB_DOCKER_CONTAINER} -p 5432:5432 -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -d postgres:12-alpine
	
create_db:
	docker exec -it ${DB_DOCKER_CONTAINER} createdb --username=${USER} --owner=${USER} ${DB_NAME}

start_container:
	docker start ${DB_DOCKER_CONTAINER}

create_migration:
	sqlx migrate add -r init 

migrate_up:
	sqlx migrate run --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"



# i was struggling with running the make file i was getting an error "make: *** No rule to make target `stop_containers'.  Stop." turns out it was because i name my makefile "MakeFile" and not "Makefile"

# docker run --name -e  -p 5432:5432 -e POSTGRES_USER=ibukunoluwaakintobi -e POSTGRES_PASSWORD= -d postgres:12-alpine

# -- the command was run and executed like this even though there is clearly no name and password showing (it couldn't access those) .env variables hen imput -e behind the name as in --name -e ${DB_DOCKER_CONTAINER}

# i figured out the issue - i commented out the include .env line wtf!

# https://stackoverflow.com/questions/48522615/docker-error-invalid-reference-format-repository-name-must-be-lowercase

# up - upload the migrations
# down - downgrade / revert the migrations 

# added a migration script 

# what is migration 