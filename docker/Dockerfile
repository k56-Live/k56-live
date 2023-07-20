# Dockerfile for content_aggregation
FROM golang:1.17

WORKDIR /go/src/github.com/k56/content_aggregation
COPY . .

RUN go mod download
RUN go build -o app main.go

CMD ["./app"]
