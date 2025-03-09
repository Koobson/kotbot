FROM golang:1.24.1-alpine3.21 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux 
RUN go build -o kotbot .

FROM alpine:3.21
ARG COMMIT_HASH="dev"
ENV COMMIT_HASH=$COMMIT_HASH
COPY --from=builder /app/kotbot .
CMD ["/kotbot"]