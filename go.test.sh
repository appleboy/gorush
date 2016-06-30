#!/usr/bin/env bash

set -e

output() {
  printf "\033[32m"
  echo $1
  printf "\033[0m"
  exit 0
}

coverage_mode=$1

test -z $coverage_mode && output "Usage: $0 coverage_mode"

test -f coverage.txt && rm -rf coverage.txt
echo "mode: ${coverage_mode}" > coverage.txt

for d in ./storage/boltdb/... ./storage/redis/... ./storage/memory/... ./config/... ./gorush/...; do
  go test -v -cover -coverprofile=profile.out -covermode=${coverage_mode} $d
  if [ -f profile.out ]; then
    sed '1d' profile.out >> coverage.txt
    rm profile.out
  fi
done
