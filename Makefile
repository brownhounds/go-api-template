run:
	@./scripts/build-docs.sh
	@go run main.go

run-compose:
	@./scripts/build-docs.sh
	docker compose up

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

dev-db:
	@env $(cat .env) docker compose up -d postgres

install-migrator:
	@sh -c "$$(curl -fsSL https://raw.githubusercontent.com/brownhounds/migoro/v0.1.3/tools/install-linux-amd64.sh)"
