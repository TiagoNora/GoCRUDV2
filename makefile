.PHONY: default run build test docs clean dev
# Variables
APP_NAME=gocrud

# Tasks
default: dev

run:
	@go run main.go
run-with-docs:
	@swag init
	@go run main.go
build:
	@go build -o $(APP_NAME) main.go
test:
	@go test ./ ...
docs:
	@swag init
clean:
	@rm -f $(APP_NAME)
	@rm -rf ./docs
dev:
	@go run main.go