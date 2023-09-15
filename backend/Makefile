include .env
export

init:
	cp -r data /tmp/
	chmod -R 777 /tmp/data

tidy:
	go mod tidy

migrate-up:
	migrate -path ../db/migrations -database $(POSTGRES_DSN) up

migrate-down:
	migrate -path ../db/migrations -database $(POSTGRES_DSN) down

backend:
	go run cmd/apg6/main.go