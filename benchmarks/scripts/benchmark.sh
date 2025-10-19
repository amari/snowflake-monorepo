#!/bin/bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)"

GHZ_CONFIG_DIR="$ROOT_DIR/benchmarks/ghz"
REPORTS_DIR="$ROOT_DIR/benchmarks/reports"

mkdir -p "$REPORTS_DIR"

echo "Running NextSnowflake benchmark..."
ghz --config "$GHZ_CONFIG_DIR/next_snowflake.toml" --output "$REPORTS_DIR/next_snowflake_report.log" --format summary

echo "Running BatchNextSnowflake benchmark..."
ghz --config "$GHZ_CONFIG_DIR/batch_next_snowflake.toml" --output "$REPORTS_DIR/batch_next_snowflake_report.log" --format summary

echo "Benchmarks complete. Reports saved to $REPORTS_DIR"
