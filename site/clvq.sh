#!/bin/sh
exec go run ./cmd/clvq-site -config ./site/clvq.json "$@"
