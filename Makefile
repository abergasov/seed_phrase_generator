BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
FILE_HASH?=$(git rev-parse HEAD)

build:
	@echo "-- building binary"
	go build -ldflags "-X main.appName=seed_generator -X main.buildHash=${FILE_HASH} -X main.buildTime=${BUILD_TIME}" -o ./bin/seed_gen ./cmd

test:
	@echo "-- testing internal modules"
	go test -race ./internal...

lint:
	@echo "-- linter running"
	golangci-lint run -c .golangci.yaml ./...