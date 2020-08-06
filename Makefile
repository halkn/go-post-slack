lint:
	@staticcheck ./...
test:
	@go test ./... -v
build:
	@go build -o bin/
clean:
	@rm -f ./bin/*
