include .env
export

all: greeting daemon client

greeting:
	@echo "Constructing Project..."

daemon:
	go build ./cmd/foodd
	export 
client:
	go build ./cmd/food