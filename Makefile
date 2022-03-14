install:
		go get .
swag:
		swag init
setup:
		make swag
		go build
run:
		make setup
	  go run key-rate-api
run-dev:
		go run main.go
