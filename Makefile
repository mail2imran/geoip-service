run:
	go run ./cmd/main.go

build:
	go build -o geoip-service ./cmd/main.go

download-mmdb:
	bash ./scripts/update_mmdb.sh $(if $(ENV_FILE),$(ENV_FILE),config/.env.dev)

serve:
	PORT=8080 ./geoip-service

test:
	go test ./... -v