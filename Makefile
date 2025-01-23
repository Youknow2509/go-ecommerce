# List Variables Path
SERVER_MAIN = ./cmd/server/main.go
WIRE_DIR = internal/wire

# List Variables Color
GREEN_COLOR_BG = \033[42m
RED_COLOR_BG = \033[41m
YELLOW_COLOR_BG = \033[43m
RESET_COLOR = \033[0m

# List Variables Wire
GO = go
WIRE = wire

# List Variables Goose
GOOSE=goose
GOOSE_NAME ?= GO_ECOMMERCE
GOOSE_DRIVER ?= mysql
GOOSE_DB_DSN = "root:root123@tcp(127.0.0.1:3306)/go_ecommerce"
GOOSE_MIGRATION_DSN ?= sql/schema
GOOSE_PATH_SCHEMA ?= sql/schema
GOOSE_PATH_QUERIES = sql/queries

# Phony Targets
.PHONY: help
.PHONY: install_path_pkg run_server wire regenerate_wire clear_log cre_env deps build test coverage
.PHONY: docker_build docker_run docker_stop docker_stop_v
.PHONY: exec_mysql mysql_dump exec_redis exec_kafka_ui
.PHONY: goose_create goose_up goose_down goose_status goose_fix goose_redo goose_reset goose_clean
.PHONY: swag

# Help Command
help:
	@echo "${GREEN_COLOR_BG}Usage: make [command]${RESET_COLOR}"
	@echo "Commands:"
	@echo "\t ${YELLOW_COLOR_BG}install_path_pkg${RESET_COLOR} \t Add the Go binary directory to your PATH"
	@echo "\t ${YELLOW_COLOR_BG}run_server${RESET_COLOR} \t Run server in development mode"
	@echo "\t ${YELLOW_COLOR_BG}wire${RESET_COLOR} \t Generate Wire dependencies"
	@echo "\t ${YELLOW_COLOR_BG}regenerate_wire${RESET_COLOR} \t Force regenerate Wire dependencies"
	@echo "\t ${YELLOW_COLOR_BG}clear_log${RESET_COLOR} \t Clear log files"
	@echo "\t ${YELLOW_COLOR_BG}cre_env${RESET_COLOR} \t Create .env from .yaml"
	@echo "\t ${YELLOW_COLOR_BG}deps${RESET_COLOR} \t Download dependencies"
	@echo "\t ${YELLOW_COLOR_BG}build${RESET_COLOR} \t Build server"
	@echo "\t ${YELLOW_COLOR_BG}test${RESET_COLOR} \t Run tests"
	@echo "\t ${YELLOW_COLOR_BG}coverage${RESET_COLOR} \t Generate coverage report"
	@echo "\nDocker Commands:"
	@echo "\t ${YELLOW_COLOR_BG}docker_build${RESET_COLOR} \t Build Docker containers"
	@echo "\t ${YELLOW_COLOR_BG}docker_run${RESET_COLOR} \t Run Docker containers"
	@echo "\t ${YELLOW_COLOR_BG}docker_stop${RESET_COLOR} \t Stop Docker containers"
	@echo "\t ${YELLOW_COLOR_BG}docker_stop_v${RESET_COLOR} \t Stop and remove Docker volumes"
	@echo "\nContainer Exec Commands:"
	@echo "\t ${YELLOW_COLOR_BG}exec_mysql${RESET_COLOR} \t Execute MySQL CLI"
	@echo "\t ${YELLOW_COLOR_BG}mysql_dump${RESET_COLOR} \t Render SQL file in database"
	@echo "\t ${YELLOW_COLOR_BG}exec_redis${RESET_COLOR} \t Execute Redis CLI"
	@echo "\t ${YELLOW_COLOR_BG}exec_kafka_ui${RESET_COLOR} \t Open Kafka UI"
	@echo "\nGoose Migration Commands:"
	@echo "\t ${YELLOW_COLOR_BG}goose_create${RESET_COLOR} \t Create a new migration"
	@echo "\t ${YELLOW_COLOR_BG}goose_up${RESET_COLOR} \t Run all available migrations"
	@echo "\t ${YELLOW_COLOR_BG}goose_down${RESET_COLOR} \t Rollback the most recent migration"
	@echo "\t ${YELLOW_COLOR_BG}goose_status${RESET_COLOR} \t Show the status of all migrations"
	@echo "\t ${YELLOW_COLOR_BG}goose_fix${RESET_COLOR} \t Fix the last migration"
	@echo "\t ${YELLOW_COLOR_BG}goose_redo${RESET_COLOR} \t Rollback and re-run the most recent migration"
	@echo "\t ${YELLOW_COLOR_BG}goose_reset${RESET_COLOR} \t Rollback all migrations"
	@echo "\t ${YELLOW_COLOR_BG}goose_clean${RESET_COLOR} \t Remove all migrations"
	@echo "\nSQLC Commands:"
	@echo "\t ${YELLOW_COLOR_BG}sqlc_generate${RESET_COLOR} \t Generate SQLC queries"
	@echo "\nSwagger Commands:"
	@echo "\t ${YELLOW_COLOR_BG}swag${RESET_COLOR} \t Handle swagger"

# SQLC - Generate
sqlc_generate:
	@echo "${YELLOW_COLOR_BG}Generating SQLC queries${RESET_COLOR}"
	sqlc generate
	@echo "${GREEN_COLOR_BG}SQLC queries generated${RESET_COLOR}"

# Goosee - Create a new migration
goose_create:
	@echo "${YELLOW_COLOR_BG}Creating a new migration${RESET_COLOR}"
	${GOOSE} -dir ${GOOSE_PATH_SCHEMA} create $(NAME) sql
	@echo "${GREEN_COLOR_BG}Migration created${RESET_COLOR}"

# Goose - Run all available migrations
goose_up:
	@echo "${YELLOW_COLOR_BG}Running all available migrations${RESET_COLOR}"
	goose -dir ${GOOSE_PATH_SCHEMA} ${GOOSE_DRIVER} $(GOOSE_DB_DSN) up
	@echo "${GREEN_COLOR_BG}Migrations completed${RESET_COLOR}"

# Goose - Rollback the most recent migration
goose_down:
	@echo "${YELLOW_COLOR_BG}Rolling back the most recent migration${RESET_COLOR}"
	goose -dir ${GOOSE_PATH_SCHEMA} ${GOOSE_DRIVER} $(GOOSE_DB_DSN) down
	@echo "${GREEN_COLOR_BG}Migration rolled back${RESET_COLOR}"

# Goose - Show the status of all migrations
goose_status:
	@echo "${YELLOW_COLOR_BG}Showing the status of all migrations${RESET_COLOR}"
	goose -dir ${GOOSE_PATH_SCHEMA} ${GOOSE_DRIVER} $(GOOSE_DB_DSN) status
	@echo "${GREEN_COLOR_BG}Migration status displayed${RESET_COLOR}"

# Goose - Fix the last migration
goose_fix:
	@echo "${YELLOW_COLOR_BG}Fixing the last migration${RESET_COLOR}"
	goose -dir ${GOOSE_PATH_SCHEMA} ${GOOSE_DRIVER} $(GOOSE_DB_DSN) fix
	@echo "${GREEN_COLOR_BG}Migration fixed${RESET_COLOR}"

# Goose - Rollback and re-run the most recent migration
goose_redo:
	@echo "${YELLOW_COLOR_BG}Rolling back and re-running the most recent migration${RESET_COLOR}"
	goose -dir ${GOOSE_PATH_SCHEMA} ${GOOSE_DRIVER} $(GOOSE_DB_DSN) redo
	@echo "${GREEN_COLOR_BG}Migration redone${RESET_COLOR}"

# Goose - Rollback all migrations
goose_reset:
	@echo "${YELLOW_COLOR_BG}Rolling back all migrations${RESET_COLOR}"
	goose -dir ${GOOSE_PATH_SCHEMA} ${GOOSE_DRIVER} $(GOOSE_DB_DSN) reset
	@echo "${GREEN_COLOR_BG}Migrations reset${RESET_COLOR}"

# Goose - Remove all migrations
goose_clean:
	@echo "${YELLOW_COLOR_BG}Removing all migrations${RESET_COLOR}"
	goose -dir ${GOOSE_PATH_SCHEMA} ${GOOSE_DRIVER} $(GOOSE_DB_DSN) clean
	@echo "${GREEN_COLOR_BG}Migrations removed${RESET_COLOR}"

# Wire Generation
wire:
	@echo "${YELLOW_COLOR_BG}Checking and generating Wire dependencies...${RESET_COLOR}"
	@if [ ! -f $(WIRE_DIR)/wire_gen.go ]; then \
		cd $(WIRE_DIR) && $(WIRE) gen; \
	fi

# Force Wire Regeneration
regenerate_wire:
	@echo "${YELLOW_COLOR_BG}Force regenerating Wire dependencies...${RESET_COLOR}"
	cd $(WIRE_DIR) && $(WIRE) gen
	@echo "${GREEN_COLOR_BG}Wire generation complete${RESET_COLOR}"

# Run Server with Wire Generation
run_server: wire
	@echo "${GREEN_COLOR_BG}Running server in development mode${RESET_COLOR}"
	@echo "${YELLOW_COLOR_BG}Configuration: config/local.yaml${RESET_COLOR}"
	$(GO) run $(SERVER_MAIN)

# Clear Log Files
clear_log:
	@echo "${YELLOW_COLOR_BG}Clearing log files${RESET_COLOR}"
	rm -rf ./storages/log/*
	rm -rf ./storages/logs/*
	@echo "${GREEN_COLOR_BG}Log files cleared${RESET_COLOR}"

# Create Environment Variables
cre_env:
	@echo "${YELLOW_COLOR_BG}Creating .env from .yaml${RESET_COLOR}"
	rm -rf .env
	$(GO) run cmd/cli/viper/main.vipper.convert.go
	@echo "${GREEN_COLOR_BG}.env file created${RESET_COLOR}"

# Docker Build
docker_build:
	@echo "${YELLOW_COLOR_BG}Building Docker containers${RESET_COLOR}"
	docker compose build
	@echo "${GREEN_COLOR_BG}Docker build complete${RESET_COLOR}"

# Docker Run
docker_run:
	@echo "${YELLOW_COLOR_BG}Running Docker containers${RESET_COLOR}"
	docker compose up -d
	@echo "${GREEN_COLOR_BG}Docker containers started${RESET_COLOR}"

# Docker Stop
docker_stop:
	@echo "${YELLOW_COLOR_BG}Stopping Docker containers${RESET_COLOR}"
	docker compose down
	@echo "${GREEN_COLOR_BG}Docker containers stopped${RESET_COLOR}"

# Docker Stop with Volume Removal
docker_stop_v:
	@echo "${YELLOW_COLOR_BG}Stopping and removing Docker volumes${RESET_COLOR}"
	docker compose down --volumes --remove-orphans
	@echo "${GREEN_COLOR_BG}Docker containers and volumes removed${RESET_COLOR}"

# MySQL Container Exec
exec_mysql:
	@echo "${YELLOW_COLOR_BG}Executing MySQL CLI${RESET_COLOR}"
	docker exec -it mysql_v8_container mysql -u root -p

# Redis Container Exec
exec_redis:
	@echo "${YELLOW_COLOR_BG}Executing Redis CLI${RESET_COLOR}"
	docker exec -it redis_v7_container redis-cli

# Kafka UI
exec_kafka_ui:
	@echo "${YELLOW_COLOR_BG}Opening Kafka UI${RESET_COLOR}"
	open http://localhost:8083

# Install Dependencies
deps:
	$(GO) mod tidy
	$(GO) mod download
	@echo "${GREEN_COLOR_BG}Dependencies installed${RESET_COLOR}"

# Build Server
build: wire
	@echo "${GREEN_COLOR_BG}Building server${RESET_COLOR}"
	$(GO) build -o server $(SERVER_MAIN)
	@echo "${GREEN_COLOR_BG}Build complete${RESET_COLOR}"

# Run Tests
test:
	@echo "${YELLOW_COLOR_BG}Running tests${RESET_COLOR}"
	$(GO) test ./... -v
	@echo "${GREEN_COLOR_BG}Tests completed${RESET_COLOR}"

# Generate Coverage Report
coverage:
	@echo "${YELLOW_COLOR_BG}Generating coverage report${RESET_COLOR}"
	$(GO) test ./... -coverprofile=coverage.out
	$(GO) tool cover -html=coverage.out
	@echo "${GREEN_COLOR_BG}Coverage report generated${RESET_COLOR}"

# Render my sql file with mysqldump 
mysql_dump:
	@echo "${YELLOW_COLOR_BG}Renaming SQL file in database${RESET_COLOR}"
	docker exec -i mysql_v8_container mysqldump -uroot --databases go_ecommerce --add-drop-database --add-drop-table --add-drop-trigger --add-locks --no-data > migrations/go_ecommerce.sql -p
	@echo "${GREEN_COLOR_BG}SQL file rendered${RESET_COLOR}"

# Add Go Binary Directory to PATH
install_path_pkg:
	@echo "${YELLOW_COLOR_BG}Adding Go binary directory to PATH${RESET_COLOR}"
	export PATH=$PATH:$(go env GOPATH)/bin
	@echo "${GREEN_COLOR_BG}Go binary directory added to PATH${RESET_COLOR}"

# handle swagger
swag:
	swag init -g ./cmd/server/main.go -o ./cmd/swag/docs