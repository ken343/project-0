include .env

all: greeting daemon client

greeting:
	@echo "Constructing Project..."

daemon:
	go build ./cmd/foodd

client:
	go build ./cmd/food