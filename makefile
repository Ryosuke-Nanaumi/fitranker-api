DATABASE_URL = $(shell grep DATABASE_URL .env | cut -d '=' -f2-)
MIGRATIONS_DIR = ./db/migrations

.PHONY: run migrate-up migrate-down seed

run:
	go run ./cmd/server

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DATABASE_URL)" down 1