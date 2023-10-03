BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} main.go

run:
	go run main.go

clean:
	go clean
	rm -f ${BINARY_NAME}
