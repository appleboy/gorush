#!/bin/sh

OS="darwin linux"
ARCH="amd64"

for GOOS in $OS; do
  for GOARCH in $ARCH; do
    EXE="gorush"
    (test "$GOOS" = "windows") && EXE="gorush.exe"

    echo "Build: ${GOOS}, Arch: ${GOARCH}, EXE: ${EXE}"
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-w" -o bin/$GOOS/$GOARCH/${EXE} gorush.go;
    tar -C bin/$GOOS/$GOARCH -czf bin/gorush-$GOOS-$GOARCH.tar.gz gorush
  done
done
