#!/bin/sh
for GOOS in darwin linux windows; do
  for GOARCH in 386 amd64; do
    EXE="gopush"
    if [ $GOOS == "windows" ]; then
      EXE="gopush.exe"
    fi

    echo "Build: ${GOOS}, Arch: ${GOARCH}, EXE: ${EXE}"
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="-w" -o bin/$GOOS/$GOARCH/${EXE} gopush.go;
  done
done
