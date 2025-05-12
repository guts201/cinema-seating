# Use the official Golang image as the build environment
FROM golang:1.23 AS build

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server cmd/main.go

FROM alpine:3.15
WORKDIR /app
COPY --from=build /app/server /app/server

COPY configs/config.yaml /app/config.yaml

EXPOSE 8080

CMD ["/app/server", "-c", "/app/config.yaml"]
