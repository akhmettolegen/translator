include .env
export

LOCAL_BIN:=$(CURDIR)/bin
PATH:=$(LOCAL_BIN):$(PATH)

help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

run: ### swag run
	go mod tidy && go mod download && \
	go run -tags migrate ./cmd/app

migrate-up: ### migration up
	migrate -path migrations -database '$(PG_URL)?sslmode=disable' up

swag-v1: ### swag init
	swag init -g internal/app/app.go