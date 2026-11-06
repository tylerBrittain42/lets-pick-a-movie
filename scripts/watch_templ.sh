#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."

go tool templ generate --watch --proxy="http://localhost:8080" --cmd="go run ."
