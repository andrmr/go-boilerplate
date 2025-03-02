@PHONY: deps build check fix test test-fast test-cov docker-prd

BUILD_OUT ?= "../.build/"
IMG_TAG ?= "golnib:latest"

_ensure-linter:
	@( \
		VERSION="1.64.5"; \
		version=$$(golangci-lint --version || echo "golangci-lint not installed"; exit 1); \
		if ! echo $$version | grep -q $$VERSION; then \
			echo "golangci-lint version is not $$VERSION"; exit 1; fi \
	) \

deps:
	@cd ./src; \
	go mod download

build:
	@cd ./src; if [ ! -d $(BUILD_OUT) ]; then mkdir $(BUILD_OUT); fi; \
	go build -o $(BUILD_OUT) ./cmd/golnib

check: _ensure-linter
	@cd ./src; \
	go mod tidy -v -diff && \
	golangci-lint -c .golangci.yaml run ./...

fix: _ensure-linter
	@cd ./src; \
	go mod tidy -v && \
	golangci-lint -c .golangci.yaml run --fix ./...

test:
	@cd ./src; \
	go test -v -race -vet=all -coverprofile=./coverage.out ./...

test-fast:
	@cd ./src; \
	go test -failfast ./...

test-cov: test
	@cd ./src; \
	go tool cover -func ./coverage.out

docker-prd:
	docker build -t $(IMG_TAG) -f ./docker/prd/Dockerfile .
