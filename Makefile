db gen:
	@sqlc generate --file sqlc/sqlc.yaml

build:
	@go build -o bin/tutorme

run:build
	@./bin/tutorme

docker:up
	@docker compose up --build
