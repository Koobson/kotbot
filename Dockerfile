FROM golang:1.24.0-alpine3.21 AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
RUN go build -o kotbot

FROM alpine:3.21
COPY --from=builder /app/kotbot .
CMD ["/kotbot"]