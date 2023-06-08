PROJECT_NAME := books-api-server
build:
	go build -o bin/${PROJECT_NAME} .
run:
	go run ./main.go
test:
	go test -v ./...
generatemocks:
	rm -rf mocks/* && \
	mockery --all --keeptree --inpackage;