.PHONY: build
build:
	@echo "build:git version=$(shell git rev-parse --default HEAD)"
	@echo "build:build protocol"
	@$(MAKE) -C protocol build_protocol >/dev/null 2>&1
	@echo "build:clean up build cache"
	@go clean $(PWD)/... >/dev/null 2>&1;true
	@echo "build:go mod download"
	@go mod download
	@echo "build:go build"
	@go build -v ./...
	@echo "build:finish"

.PHONY: clean
clean:
	@echo "clean:clean up protocol"
	@$(MAKE) -C protocol clean >/dev/null 2>&1
	@echo "clean:clean up build cache"
	@go clean . >/dev/null 2>&1;true