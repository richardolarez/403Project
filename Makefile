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
	go clean &
	npm --prefix web/app cache clean --force
	rm -f ${BINARY_NAME}

deploy:
	go run main.go &
	cd web/app && npm install &
	npm --prefix web/app run build &
	

