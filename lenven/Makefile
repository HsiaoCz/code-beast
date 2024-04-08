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

proto:
	@protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	pb/lenven.proto