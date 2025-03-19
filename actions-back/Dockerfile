# Stage 1: Build
FROM golang:1.24 AS builder

WORKDIR /app-ms

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main-ms ./cmd/main.go

# Stage 2: Image
FROM alpine:latest

WORKDIR /app-ms

RUN apk add --no-cache ca-certificates

COPY --from=builder /app-ms/main-ms .

EXPOSE 8080

CMD ["./main-ms"]
