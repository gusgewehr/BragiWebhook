FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o BragiWebhook .

FROM debian:bookworm

WORKDIR /app

COPY --from=builder /app/BragiWebhook .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./BragiWebhook"]