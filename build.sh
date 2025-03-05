#!/usr/bin/env bash
set -euo pipefail

cd frontend
#pnpm install --frozen-lockfile
#pnpm run build

cd ..
GIT_COMMIT=$(git log -n 1 --format=%h)
VERSION=$(git describe --tags --abbrev=0 --match=v* | cut -c 2-)

LDFLAGS="-s -w -X github.com/filebrowser/filebrowser/v2/version.Version=$VERSION -X github.com/filebrowser/filebrowser/v2/version.CommitSHA=$GIT_COMMIT"
go build -ldflags="$LDFLAGS" -o filebrowser .
