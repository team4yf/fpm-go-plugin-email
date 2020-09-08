PROJECTNAME=$(shell basename "$(PWD)")

GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

all: build

install:
	go mod download

dev:
	docker-compose -f docker-compose-dev.yml up -d

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o $(GOBIN)/app ./main.go || exit

start:
	go build -o $(GOBIN)/app ./main.go || exit
	$(GOBIN)/app