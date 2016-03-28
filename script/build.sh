#!/bin/sh

OS="darwin linux windows"
ARCH="amd64"

for GOOS in $OS; do
  for GOARCH in $ARCH; do
    EXE="gopush"
    (test "$GOOS" = "windows") && EXE="gopush.exe"

    echo "Build: ${GOOS}, Arch: ${GOARCH}, EXE: ${EXE}"
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-w" -o bin/$GOOS/$GOARCH/${EXE} gopush.go;
  done
done
