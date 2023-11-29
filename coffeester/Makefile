




# include .env


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


# i was struggling with running the make file i was getting an error "make: *** No rule to make target `stop_containers'.  Stop." turns out it was because i name my makefile "MakeFile" and not "Makefile"

