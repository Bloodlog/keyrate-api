build:
		go mod download
run:build
		go run cmd/app/main.go
swag:
		swag init -g cmd/app/main.go
test:
	  go test -v ./...
