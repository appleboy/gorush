@ECHO OFF
FOR /F "tokens=* USEBACKQ" %%F IN (`git describe --tags --always`) DO (
SET VERSION=%%F
)
echo go build -v -ldflags '-s -w -X main.Version=%VERSION%' -o bin
go build -v -ldflags '-w' -o bin/gorush.exe