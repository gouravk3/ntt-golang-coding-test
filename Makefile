.PHONY: all build run stop clean

DOCKER_COMPOSE_FILE=docker-compose.yaml

all: build

build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

run:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

logs:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

stop:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

clean: stop
	docker-compose -f $(DOCKER_COMPOSE_FILE) down --rmi all
