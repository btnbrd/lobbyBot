FROM golang:1.24-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bot ./cmd/bot

FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root/

COPY --from=build /app/bot .

# debug output - remove later
#ENV TELEGRAM_TOKEN=debug-missing-token

CMD ["./bot"]
