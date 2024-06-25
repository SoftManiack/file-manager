FROM golang:alpine
COPY go.mod .
RUN go mod download
RUN go run cmd/main.go