# Build environment
# -----------------
FROM golang:1.18-alpine as build-env
WORKDIR /keyrate-api

RUN apk update && apk add --no-cache gcc musl-dev git

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -ldflags '-w -s' -a -o ./bin/app ./cmd/app


# Deployment environment
# ----------------------
FROM alpine
RUN apk update

COPY --from=build-env /keyrate-api/bin/app /keyrate-api/

EXPOSE 8080
CMD ["/myapp/app"]
