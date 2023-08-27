# Build stage
FROM golang:1.20.1-alpine3.16 AS builder
WORKDIR /app

COPY . .
RUN go build -o main ./cmd/server/main.go

# Run stage
FROM alpine:3.16
WORKDIR /app

COPY --from=builder /app/main .
COPY ./cmd/server/app.env .

EXPOSE 8080
CMD [ "/app/main" ]

