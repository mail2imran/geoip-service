# 🌍 GeoIP Country Check Service

A Golang-based API service to validate if an IP address belongs to a whitelisted country using MaxMind GeoLite2 data. This is useful for preventing outsourced or out-of-region access in sensitive applications.

---

## Features

- IP address to country validation using MaxMind GeoLite2 DB
- JWT-based authentication for secure access
- Docker support for containerized deployment
- Kubernetes CronJob to keep GeoIP data up-to-date
- Environment-based config for dev/staging/prod
- Unit and integration tests with test `.mmdb` support

---

## Project Structure

```
geoip-service/
├── cmd/                         # Main application entrypoint
├── config/                      # .env files and downloaded .mmdb
├── internal/
│   ├── geoip/                   # MaxMind resolver
│   ├── handler/                 # API endpoints
│   └── middleware/              # JWT middleware
├── scripts/                     # DB download script
├── k8s/                         # Kubernetes CronJob manifest
├── Dockerfile                   # Container definition
├── Makefile                     # Run/test/build helpers
└── README.md
```

---

## ⚙️ Setup

1. **Clone the repo**
2. **Install Go 1.23+**
3. **Download the GeoIP database**

```bash
make download-mmdb
```

> Optional: specify lifecycle environment

```bash
make download-mmdb ENV_FILE=config/.env.prod
```

4. **Run the app**

```bash
make run
```

5. **Test the endpoints**

Login:
```bash
curl -X POST http://localhost:8080/auth/token \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'
```

Check IP:
```bash
curl -X POST http://localhost:8080/ip/verify \
  -H "Authorization: Bearer <TOKEN>" \
  -H "Content-Type: application/json" \
  -d '{"ip":"8.8.8.8", "allowed_countries": ["US"]}'
```

---

## MaxMind DB Update Plan

### Local
Run the included shell script:

```bash
make download-mmdb
```

- Pulls the latest `.mmdb` using your `MAXMIND_LICENSE_KEY`
- Extracts and places it into `config/GeoLite2-Country.mmdb`

### Production (Kubernetes)
A scheduled CronJob (`k8s/cronjob-update-mmdb.yaml`) updates the DB weekly:

- Mounts shared volume (`PVC`) with your app
- Downloads and replaces the DB using a secret-stored license key

> Be sure to create the secret first:

```bash
kubectl create secret generic maxmind-secret \
  --from-literal=license-key=your_key_here
```

---

## Testing

Ensure test `.mmdb` exists:

```bash
cp config/GeoLite2-Country.mmdb internal/geoip/testdata/GeoLite2-Country-Test.mmdb
make test
```

---

## Development Notes

- `.env` files per environment go inside `config/`
- Uncommitted secrets and `.mmdb` are ignored via `.gitignore`

---

## License

This project uses the [MaxMind GeoLite2](https://dev.maxmind.com/geoip/geolite2/) data under its free data license.
