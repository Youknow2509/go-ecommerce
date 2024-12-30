# List Variables Color
GREEN_COLOR_BG = \033[42m
RED_COLOR_BG = \033[41m
YELLOW_COLOR_BG = \033[43m
RESET_COLOR = \033[0m

# ----------------------------------------------------------------
.PHONY: help run_server clear_log

help:
	@echo "${GREEN_COLOR_BG}usage: make [command]${RESET_COLOR}"
	@echo "list commands:"
	@echo "\t ${YELLOW_COLOR_BG}run_server${RESET_COLOR} \t Run the server in development mode"
	@echo "\t ${YELLOW_COLOR_BG}clear_log${RESET_COLOR} \t Clear the log file"
	@echo "\t ${YELLOW_COLOR_BG}docker_build${RESET_COLOR} \t Build all container docker use"
	@echo "\t ${YELLOW_COLOR_BG}docker_run${RESET_COLOR} \t Run all container docker use"
	@echo "\t ${YELLOW_COLOR_BG}docker_stop${RESET_COLOR} \t Stop all container docker use"
	@echo "\t ${YELLOW_COLOR_BG}docker_stop_v${RESET_COLOR} \t Docker stop and remove volumes"

run_server:	
	@echo "run server in development mode (can change to production mode in file config .yaml) \e[0m"
	go run ./cmd/server/main.go

clear_log:
	@echo "${YELLOW_COLOR_BG}clear all file loggers${RESET_COLOR}"
	rm -rf ./storages/log/*
	rm -rf ./storages/logs/*

docker_build:
	@echo "${YELLOW_COLOR_BG}docker build all container project${RESET_COLOR}"
	docker compose build

docker_run:
	@echo "${YELLOW_COLOR_BG}docker run all container project${RESET_COLOR}"
	docker compose up -d

docker_stop:
	@echo "${YELLOW_COLOR_BG}docker stop all container project${RESET_COLOR}"
	docker compose down

docker_stop_v:
	@echo "${YELLOW_COLOR_BG}docker stop, remove volumes and networks from project${RESET_COLOR}"
	docker compose down --volumes --remove-orphans

