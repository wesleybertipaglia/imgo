PROJECT_NAME = imgool
BIN_DIR = bin
OUTPUT_DIR = output
SRC_DIR = src

GO = go
GO_BUILD = $(GO) build
GO_RUN = $(GO) run

GO_FILES = $(SRC_DIR)/*.go

dev: 
	$(GO_RUN) main.go

build: 
	$(GO_BUILD) -o $(PROJECT_NAME)

clean:
	rm -f $(PROJECT_NAME)

install:
	sudo mv $(PROJECT_NAME) /usr/local/bin/

start:
	$(GO_RUN) main.go

fmt:
	$(GO) fmt $(GO_FILES)

lint:
	# Example: Install and run a linter (optional)
	golangci-lint run

test:
	$(GO) test ./...

help:
	@echo "Makefile Commands:"
	@echo "  make dev      	- Runs the application with 'go run main.go'"
	@echo "  make build    	- Builds the binary executable"
	@echo "  make clean    	- Removes the built binary"
	@echo "  make install  	- Installs the binary globally"
	@echo "  make start    	- Runs the application (equivalent to 'go run main.go')"
	@echo "  make fmt      	- Formats the Go code"
	@echo "  make lint		- Lint the Go code"
