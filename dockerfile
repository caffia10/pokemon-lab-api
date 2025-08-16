FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o server ./cmd/main.go

FROM debian:bookworm-slim

RUN useradd -m appuser

WORKDIR /app
COPY --from=builder /app/server .

EXPOSE 3030

USER appuser

CMD ["./server"]
