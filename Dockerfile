FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download && go mod tidy

COPY .. ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o compra ./cmd/start

# Fase final
FROM alpine:3.20.3

WORKDIR /app

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/compra .

EXPOSE 3000

ENTRYPOINT ["./compra"]