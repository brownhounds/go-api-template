run:
	@./scripts/build-docs.sh
	@go run main.go

build:
	@GOOS=linux go build -ldflags="-s -w" -o ./bin/api main.go

lint:
	@golangci-lint run --fast

fix:
	@golangci-lint run --fix

init:
	@cp .env.example .env

install-hooks:
	@pre-commit install

docs:
	@./scripts/build-docs.sh

dev_db:
	@env $(cat .env) docker compose up -d postgres
