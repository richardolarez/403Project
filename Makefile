BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} main.go

run:
	go run main.go &
	npm --prefix web/app start

.PHONY: test
test:
	go test ./test

clean:
	go clean
	npm --prefix web/app update
	rm -f ${BINARY_NAME}
