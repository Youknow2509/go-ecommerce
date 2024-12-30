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

# Phony Targets
.PHONY: help run_server clear_log cre_env docker_build docker_run docker_stop docker_stop_v 
.PHONY: exec_mysql exec_redis exec_kafka_ui wire regenerate_wire deps build test coverage

# Help Command
help:
	@echo "${GREEN_COLOR_BG}Usage: make [command]${RESET_COLOR}"
	@echo "Commands:"
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
	@echo "\t ${YELLOW_COLOR_BG}exec_redis${RESET_COLOR} \t Execute Redis CLI"
	@echo "\t ${YELLOW_COLOR_BG}exec_kafka_ui${RESET_COLOR} \t Open Kafka UI"

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
