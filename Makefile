COMPOSE=docker-compose
COMPOSE_FILE=./compose.yml

.PHONY: up down restart

up:
	$(COMPOSE) -f $(COMPOSE_FILE) up -d
down:
	$(COMPOSE) down
restart: down up


