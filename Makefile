build:
	go mod download
run:build
	go run cmd/app/main.go
install-swag:
	go install github.com/swaggo/swag/cmd/swag@latest
swag:
	~/go/bin/swag init -g cmd/app/main.go
test:
	go test -v ./...
