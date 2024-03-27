build:
	@go build -o bin/hotel main.go

run: build
	@./bin/hotel

test:
	@go test -v ./...
	
mongo:
	@docker run --name mongodb -d mongo:latest -p 27017:27017