#!/bin/sh
set -eu
export CLVQ_ADMIN_TPL_DEV
CLVQ_ADMIN_TPL_DEV=1
exec go run ./cmd/clvq-site -config ./site/clvq.json "$@"
