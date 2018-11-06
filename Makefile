build:
	@go build main.go

run: build
	@./main -filename="gistfile1.txt"

test:
	@go get github.com/stretchr/testify/assert
	@go test ./...
