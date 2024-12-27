# Variables

# ----------------------------------------------------------------
help:
	@echo "usage: make [command]"
	@echo "list commands:"
	@echo "\t run_server \t Run the server in development mode"
	@echo "\t clear_log \t Clear the log file"

run_server:	
	go run ./cmd/server/main.go

clear_log:
	rm -rf ./log/*
	mkdir -p ./log