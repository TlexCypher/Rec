COMPOSE=docker-compose
COMPOSE_FILE=./compose.yml
GO_APP_CONTAINER=go-app
DB_CONTAINER=mysql

include ./db/.env

.PHONY: up down restart migrate-add
up:
	$(COMPOSE) -f $(COMPOSE_FILE) up -d
down:
	$(COMPOSE) down
restart: down up
migrate-add:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make create-migration NAME=<migration_name>"; \
		exit 1; \
	fi
	$(COMPOSE) exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" create $(NAME) sql
migrate-up:
	$(COMPOSE) exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" up
migrate-up-one:
	$(COMPOSE) exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" up-by-one
migrate-down:
	$(COMPOSE) exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" down 
migrate-status:
	$(COMPOSE) exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" status
migrate-version:
	$(COMPOSE) exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" version 
migrate-reset:
	$(COMPOSE) exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" down-to 0

