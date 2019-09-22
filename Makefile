lint:
	@golangci-lint run -v --enable-all
test:
	@go test ./... -v
build:
	@go build -o bin/
clean:
	@rm -f ./bin/*
