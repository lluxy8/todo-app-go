# ---------- BUILD ----------
FROM golang:1.25.6-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/api


# ---------- DEV ----------
FROM alpine:3.20 AS dev
WORKDIR /app

COPY --from=build /app/app .

EXPOSE 8080
CMD ["./app"]


# ---------- DEBUG ----------
FROM golang:1.25.6-alpine AS debug
WORKDIR /app

COPY . .
RUN go install github.com/go-delve/delve/cmd/dlv@latest

EXPOSE 8080 40000
CMD ["dlv", "debug", "cmd/api", "--headless", "--listen=:40000", "--api-version=2", "--accept-multiclient"]
