BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} main.go

run:
	go run main.go

.PHONY: test
test:
	go test ./test

clean:
	go clean
	rm -f ${BINARY_NAME}
