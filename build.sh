#!/bin/sh

# TODO: Improve this build script and include other platforms.

go build -o jiggle.exe -ldflags -H=windowsgui main.go
