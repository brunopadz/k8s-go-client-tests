FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY main.go .

RUN go build -o app .

ENTRYPOINT ["./app"]
