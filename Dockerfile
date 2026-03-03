FROM golang:1.25 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

FROM alpine:3.23

WORKDIR /app

RUN adduser -D appuser

COPY --from=builder /app/main .

USER appuser

CMD ["./main"]
