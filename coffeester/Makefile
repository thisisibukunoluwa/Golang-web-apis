

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
migrate_down:
	sqlx migrate revert --database-url "postgres://${USER}:${PASSWORD}@${HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"
build:
	if [ -f "${BINARY}" ]; then \
		rm ${BINARY}; \
		echo "Deleted ${BINARY}"; \
	fi
	@echo "Building binary"
	go build -o ${BINARY} cmd/server/*.go 

run: build
	./${BINARY}
	@echo "starting binary"

stop:
	@echo "stopping server..."
	@-pkill -SIGTERM -f "./{BINARY}"
	@echo "server stopped"

# i was struggling with running the make file i was getting an error "make: *** No rule to make target `stop_containers'.  Stop." turns out it was because i name my makefile "MakeFile" and not "Makefile"

# docker run --name -e  -p 5432:5432 -e POSTGRES_USER=ibukunoluwaakintobi -e POSTGRES_PASSWORD= -d postgres:12-alpine

# -- the command was run and executed like this even though there is clearly no name and password showing (it couldn't access those) .env variables hen imput -e behind the name as in --name -e ${DB_DOCKER_CONTAINER}

# i figured out the issue - i commented out the include .env line wtf!

# https://stackoverflow.com/questions/48522615/docker-error-invalid-reference-format-repository-name-must-be-lowercase

# up - upload the migrations
# down - downgrade / revert the migrations 

# added a migration script 

# what is migration 

# In the context of a database, migration refers to the process of evolving the database schema and data over time. Database migration is typically necessary when you need to make structural changes to the database, such as adding or removing tables, columns, or indexes, in a way that doesn't disrupt the existing data.

# There are two main types of migrations:

# Schema Migration:

# Definition: Schema migration involves changing the structure or schema of the database, such as creating or altering tables, adding or removing columns, and modifying indexes or constraints.
# Purpose: It is used to update the database schema to reflect changes in the application's data model.
# Data Migration:

# Definition: Data migration involves moving or transforming existing data to align with the new schema resulting from a schema migration.
# Purpose: It ensures that the data remains consistent and usable after structural changes to the database.

## in PgAdmin - main_coffester_db -> Databases -> Schemas -> public -> _sqlx_migrations -> 

## In Go (Golang), the underscore (_) character has a special meaning when used in an import statement. When you import a package with the underscore, it is indicating to the Go compiler that you want to use the package purely for its side effects, and you don't intend to reference any of its identifiers directly in your code.

