#!/bin/sh
for GOOS in darwin linux windows; do
  for GOARCH in 386 amd64; do
    echo "Build: ${GOOS}, Arch: ${GOARCH}"
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-w" -o bin/$GOOS/$GOARCH/gopush gopush.go;
  done
done
