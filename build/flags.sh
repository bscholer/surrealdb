#!/usr/bin/env sh

set -eu

cd "$(dirname "${0}")/.."

echo '-X "github.com/surrealdb/surrealdb/util/build.rev='$(git rev-parse HEAD)'"' \
     '-X "github.com/surrealdb/surrealdb/util/build.ver='$(git describe --tags --abbrev=0 || echo 0.0.0)'"' \
     '-X "github.com/surrealdb/surrealdb/util/build.time='$(date -u '+%Y/%m/%d %H:%M:%S')'"'
