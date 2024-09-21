build:
	@go build -o bin/janabyte cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/janabyte
