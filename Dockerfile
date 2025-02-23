FROM golang:1.24.0-bookworm AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=1
ENV GOOS=linux 
RUN go build -o kotbot -a -ldflags '-linkmode external -extldflags "-static"' .

FROM scratch
COPY --from=builder /app/kotbot .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
CMD ["/kotbot"]