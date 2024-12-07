COMPOSE=docker-compose
COMPOSE_FILE=./compose.yml
GO_APP_CONTAINER=go-app
DB_CONTAINER=mysql-db
DOCKER=docker

include ./db/.env

.PHONY: build up down restart migrate-add
build:
	$(COMPOSE) build
up:
	$(COMPOSE) -f $(COMPOSE_FILE) up -d
down:
	$(COMPOSE) down
restart: down build up
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
oapi-gen:
	$(COMPOSE) exec $(GO_APP_CONTAINER) npx -y @redocly/cli bundle openapi/openapi.yaml -o openapi/oapi-concat.yaml
	$(COMPOSE) exec $(GO_APP_CONTAINER) oapi-codegen -config ./openapi/config.yaml openapi/oapi-concat.yaml
in-db:
	$(DOCKER) container exec -it $(DB_CONTAINER) bash
gorm-gen:
	$(COMPOSE) exec $(GO_APP_CONTAINER) go run ./cmd/gorm-gen/
wire-gen:
	$(COMPOSE) exec $(GO_APP_CONTAINER) wire ./cmd/wire-gen/

