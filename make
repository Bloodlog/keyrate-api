install:
	go get .
swag:
	swag init
build: swag
	go build
run: build
  go run key-rate-api
run-dev:
	go run main.go
