# Simple Makefile for a Go project

# Build the application
build:
	@echo "Building..."
	
	
	@go build -o tasks main.go
	@sudo mv tasks /usr/local/bin/

# Clean the binary
clean:
	@echo "Cleaning..."
	@sudo rm -f /usr/local/bin/tasks


.PHONY: build clean
