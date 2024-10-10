# If the first argument is "run"...
ifeq (migrate.new,$(firstword $(MAKECMDGOALS)))
  MIG_NEW_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(MIG_NEW_ARGS):;@:)
endif

# Read .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif


BIN_DIR=bin

build:
	mkdir -p $(BIN_DIR) #
	go build -o $(BIN_DIR)/ -v ./cmd/app

tests:
	go test ./...

migrate.new:
	migrate create -ext sql -dir migrations/ -seq $(MIG_NEW_ARGS)

migrate.up:
	migrate -path ./migrations -database 'postgres://$(DATABASE_USER):$(DATABASE_PASS)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable' up

migrate.down:
	migrate -path ./migrations -database 'postgres://$(DATABASE_USER):$(DATABASE_PASS)@$(DATABASE_HOST):$(DATABASE_PORT)/$(DATABASE_NAME)?sslmode=disable' down