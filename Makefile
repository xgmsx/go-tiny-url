export

## Migrations

DB_MIGRATE_URL = postgres://login:pass@localhost:5432/app-db?sslmode=disable
MIGRATE_PATH = ./migrations

.PHONY: migrate-install
migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

.PHONY: migrate-create
migrate-create:  # usage: make migrate-create name=init
	migrate create -ext sql -dir "$(MIGRATE_PATH)" $(name)

.PHONY: migrate-up
migrate-up:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" up

.PHONY: migrate-down
migrate-down:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" down -all

## Proto

.PHONY: proto-install
proto-install:
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

.PHONY: proto-generate
proto-generate:
	protoc --go_out=generated/protobuf --go-grpc_out=generated/protobuf --proto_path=proto proto/shortener_v1.proto

## Openapi

.PHONY: openapi-generate
openapi-generate:
	swag init --generalInfo ./pkg/http/server.go --parseInternal

## Docker-compose

.PHONY: up
up: ## 🐳🔼 Start Docker containers with docker-compose
	docker-compose up -d --build

.PHONY: down
down: ## 🐳🔽 Stop Docker containers with docker-compose
	docker-compose down

## Tests, linting, generation

.PHONY: generate
generate: ## Generate artifacts
	@echo "* Running proto-generate..."
	$(MAKE) proto-generate
	@echo "* Running openapi-generate..."
	$(MAKE) openapi-generate

.PHONY: cov
cov: ## ☔ Generate a coverage report
	go test -cover -coverprofile=coverage.txt ./...
	go tool cover -html=coverage.txt -o coverage.html
	go tool cover -func=coverage.txt

.PHONY: fmt
fmt: ## 🎨 Fix code format issues
	go fmt ./...
	sort-imports ./...

.PHONY: lint
lint: ## 🚨 Run lint checks
	golangci-lint run --timeout 5m
	gocritic check -enableAll ./...

## Building, running, escape analysis

.PHONY: build
build: ## 📦 Build the program
	go build -o app.bin ./cmd/app

.PHONY: run
run: ## 🚶 Run the program
	go run ./cmd/app

.PHONY: rune
rune: ## 🔎 Run the program with escape analysis
	go run -gcflags='-m=3' ./cmd/app

## Tests, tests, tests...

.PHONY: test
test: ## 🚦 Execute unittests
	go test ./...

.PHONY: test-race
test-race: ## 🚦🏁 Execute tests with the data race detector
	go test -race -short ./...

.PHONY: test-msan
test-msan: ## 🚦🧼 Execute tests with the memory sanitizer
	go test -msan -short ./...

.PHONY: test-bench
test-bench: ## 📈 Execute benchmark tests
	go test -bench=. ./...

## Pprof

.PHONY: pprof
pprof: ## 📈 Show pprof report
	go tool pprof -http=:8001 http://localhost:8000/debug/pprof/allocs?debug=1
