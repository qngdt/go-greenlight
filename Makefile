# Include variables from the .env file
include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

.PHONY: confirm
confirm:
	@echo -n 'Are you sure? [y/N] ' && re ans && [ $${ans:-N} = y ]

## run/api: run the cmd/api application
.PHONY: run/api
run/api:
	go run ./cmd/api

.PHONY: build/api 
build/api:
	@echo 'Building cmd/api...'
	go build -ldflags='-s' -o=./bin/api ./cmd/api

## db/migrations/new name=$1: create a new database migration
.PHONY: db/migrations/new
db/migrations/new:
	@echo 'Creating migration files for ${name}...'
	dbmate new ${name}

## db/migrations/up: apply all up database migrations
.PHONY: db/migrations/up
db/migrations/up: confirm
	@echo 'Running up migrations...'
	dbmate up

.PHONY: audit
audit:
	@echo 'Tidying and verifying module dependencies...' go mod tidy
	go mod verify
	@echo 'Formatting code...'
	go fmt ./...
	@echo 'Vetting code...'
	go vet ./...
	staticcheck ./...
	@echo 'Running tests...'
	go test -race -vet=off ./...