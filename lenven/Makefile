build:
	@go build -o bin/lenven main.go

run: build
	@./bin/lenven

test:
	@go test -v ./...

gen:
	@templ generate

mongo:
	@docker run 