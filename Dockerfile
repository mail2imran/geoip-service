FROM golang:1.23

WORKDIR /app
COPY . .
ENV ENV_PATH=config/.env.dev

RUN go mod tidy && go build -o geoip-service ./cmd/main.go

EXPOSE 8080
CMD ["./geoip-service"]