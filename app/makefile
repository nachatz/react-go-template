# Go parameters
GOCMD    := go
GOBUILD  := $(GOCMD) build
GOMOD    := $(GOCMD) mod
GOFMT    := $(GOCMD) fmt
BINARY_NAME := my_app

list: 
	echo Commands: init build run tidy fmt clean [install]

init:
	$(GOMOD) init github.com/nachatz/React-Go-Template/app

build:
	cd ./cmd && $(GOBUILD) -o ../$(BINARY_NAME)

run: tidy build
	./$(BINARY_NAME)

tidy:
	$(GOMOD) tidy

fmt:
	$(GOFMT) ./...

clean:
	@echo "Cleaning up..."
	@rm -f go.mod
	@rm -f go.sum
	@rm -f $(BINARY_NAME)

install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp $(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)

.PHONY: init build run tidy fmt clean install
