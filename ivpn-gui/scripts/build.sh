#!/usr/bin/env bash

set -uex

BINARY=ivpnc

if [ "${GOOS:-}" = "windows" ]; then
    BINARY=${BINARY}.exe
fi

CGO_ENABLED=0 go build -trimpath -ldflags="-w -s" -o build/$BINARY main.go
