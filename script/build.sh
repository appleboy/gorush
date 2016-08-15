#!/bin/sh

OS="darwin linux"
ARCH="amd64"
VERSION=$1

for GOOS in $OS; do
  for GOARCH in $ARCH; do
    EXE="gorush"
    (test "$GOOS" = "windows") && EXE="gorush.exe"

    echo "Build: ${GOOS}, Arch: ${GOARCH}, EXE: ${EXE}"
    GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -ldflags="-s -w -X main.Version=${VERSION}" -o bin/$GOOS/$GOARCH/${EXE} gorush.go;
    tar -C bin/$GOOS/$GOARCH -czf bin/gorush-$VERSION-$GOOS-$GOARCH.tar.gz gorush
  done
done
