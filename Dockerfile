FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o airline-reservation ./cmd/main.go

FROM debian:latest

WORKDIR /root/

COPY --from=builder /app/airline-reservation .

EXPOSE 8080

CMD ["./airline-reservation"]
