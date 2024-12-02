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
mysql "user:password@/dbname?parseTime=true"
migrate-add:
	@if [ -z "$(NAME)" ]; then \
		echo "Error: NAME is required. Usage: make create-migration NAME=<migration_name>"; \
		exit 1; \
	fi
migrate-up:
	docker-compose exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" up
migrate-up-one:
	docker-compose exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" up-by-one
migrate-down:
	docker-compose exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" down 
migrate-status:
	docker-compose exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" status
migrate-version:
	docker-compose exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" version 
migrate-reset:
	docker-compose exec $(GO_APP_CONTAINER) goose -dir=./migrations mysql "$(MYSQL_USER):$(MYSQL_PASSWORD)@tcp(mysql:$(MYSQL_PORT))/$(MYSQL_DATABASE)" down-to 0

