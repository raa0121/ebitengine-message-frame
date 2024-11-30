@echo off

set GOOS=js
set GOARCH=wasm

go build -o docs\message-frame.wasm .\example

set GOOS=
set GOARCH=
