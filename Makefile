build:
	@go build -o bin/backend cmd/server/main.go

test:
	@go test -v ./...

run: build
	@./bin/backend