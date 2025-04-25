#!/bin/bash
set -e

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

ENV_FILE="${1:-$ROOT_DIR/config/.env.dev}"
echo "Using env file: $ENV_FILE"

# Load .env file
if [[ ! -f "$ENV_FILE" ]]; then
  echo "Env file not found: $ENV_FILE"
  exit 1
fi

source "$ENV_FILE"

if [[ -z "$MAXMIND_LICENSE_KEY" ]]; then
  echo "MAXMIND_LICENSE_KEY is not set"
  exit 1
fi

DB_EDITION_ID="GeoLite2-Country"
DEST_DIR="$ROOT_DIR/config"
TMP_DIR="$DEST_DIR/tmp-mmdb"

echo "Downloading MaxMind DB..."
mkdir -p "$TMP_DIR"
curl -s -L -o "$TMP_DIR/db.tar.gz" "https://download.maxmind.com/app/geoip_download?edition_id=${DB_EDITION_ID}&license_key=${MAXMIND_LICENSE_KEY}&suffix=tar.gz"

echo "Extracting..."
tar -xzf "$TMP_DIR/db.tar.gz" -C "$TMP_DIR"

echo "Locating .mmdb file..."
MMDB_PATH=$(find "$TMP_DIR" -name "${DB_EDITION_ID}.mmdb" | head -n 1)

if [[ -z "$MMDB_PATH" ]]; then
  echo ".mmdb file not found after extraction"
  exit 1
fi

echo "Moving to $DEST_DIR/${DB_EDITION_ID}.mmdb"
mv "$MMDB_PATH" "$DEST_DIR/${DB_EDITION_ID}.mmdb"

echo "Cleaning up temp files"
rm -rf "$TMP_DIR"

echo "GeoIP DB updated successfully at $DEST_DIR/${DB_EDITION_ID}.mmdb"
