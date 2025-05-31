BINARY=server

.PHONY: clean build run

clean:
	@echo "Removing binaries..."
	@rm -f ./cmd/server/$(BINARY)

build:
	@echo "Building the server..."
	@go build -o ./cmd/server/$(BINARY) ./cmd/server

run: build
	@echo "Running the server..."
	@./cmd/server/$(BINARY)