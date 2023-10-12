BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} main.go

run:
	go run main.go

generate:
	java -jar swagger-codegen-cli.jar generate -i openapi.yml -l go -o server

.PHONY: test

test:
	go test ./test

clean:
	go clean
	rm -f ${BINARY_NAME}