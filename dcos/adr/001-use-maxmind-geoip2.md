# ADR 001: Use MaxMind GeoLite2 for GeoIP Country Detection

## Status

Accepted

## Context

We need a reliable and regularly updated way to determine the country of origin for incoming IP addresses in order to enforce regional access restrictions for customers.

GeoIP detection is required for:
- Blocking access from unauthorized countries
- Auditing user behavior by location
- Enforcing customer-requested regional compliance policies

Several options were considered:

- IP2Location: commercial with limited free tier
- DB-IP: permissive license but low resolution in free version
- **MaxMind GeoLite2**: well-maintained, widely used, free tier available, has a stable API for downloads

## Decision

We will use the **MaxMind GeoLite2-Country** database for GeoIP detection. This provides country-level accuracy with weekly updates, accessible via a license key.

## Consequences

### Implementation

- We'll download the `.mmdb` file manually during local development and via a **Kubernetes CronJob** in production
- We will store the license key in `.env` files locally and in a **Kubernetes secret** in production
- The database will be stored under `config/GeoLite2-Country.mmdb`
- The download script will live in `scripts/update_mmdb.sh`

### Maintenance

- MaxMind requires registering for a free account to obtain a license key
- A download limit (2,000/month) applies to the free tier
- A shell script is provided to automate download:
  ```bash
  make download-mmdb ENV_FILE=config/.env.prod
