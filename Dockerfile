FROM golang:1.18.2-alpine3.15 as builder
RUN mkdir /app
RUN mkdir /app/logs
ADD . /app
WORKDIR /app
RUN go build -o . cmd/main.go
CMD ["/app/main"]