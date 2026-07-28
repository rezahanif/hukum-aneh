#!/bin/bash
# Run the Drive→PG ingest script with the correct env vars.
# First run: opens browser for Google OAuth.
# Subsequent runs: reuses saved token.json.

set -e

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

cd "$SCRIPT_DIR"

export POSTGRES_HOST="${POSTGRES_HOST:-localhost}"
export POSTGRES_PORT="${POSTGRES_PORT:-5432}"
export POSTGRES_DB="${POSTGRES_DB:-hukum_aneh}"
export POSTGRES_USER="${POSTGRES_USER:-hukum}"
export POSTGRES_PASSWORD="bismillahip3.5bro"
export DRIVE_FOLDER_ID="1vneHF9YxwgSnBh3ashORK0cYPo16vmQS"
export DRIVE_CREDENTIALS_PATH="/home/z/my-project/upload/client_secret.json"
export DRIVE_TOKEN_PATH="/home/z/my-project/upload/token.json"

# Set TESSDATA_PREFIX for Indonesian OCR
export TESSDATA_PREFIX="${TESSDATA_PREFIX:-$HOME/.tessdata}"

echo "=== Hukum Aneh — Drive Ingest ==="
echo "PG: $POSTGRES_HOST:$POSTGRES_PORT/$POSTGRES_DB"
echo "Drive folder: $DRIVE_FOLDER_ID"
echo "Credentials: $DRIVE_CREDENTIALS_PATH"
echo "Token: $DRIVE_TOKEN_PATH"
echo ""

python3 "$SCRIPT_DIR/ingest_drive.py" "$@"