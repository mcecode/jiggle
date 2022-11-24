#!/bin/sh

cd $(realpath $(dirname "$0"))

if [ -d ./build ]; then
  echo "Deleting build directory"
  rm -r build
fi

echo "Creating build directory"
mkdir build

# OS checking logic adapted from
# https://stackoverflow.com/q/394230
# https://stackoverflow.com/q/3466166
# https://en.wikipedia.org/wiki/Uname

fynecmd="fyne package -name jiggle -icon assets/cursor.png"
os=$(uname -s | tr "[:upper:]" "[:lower:]")
case "$os" in
  "cygwin"*|"mingw"*|"msys"*|"windows"*)
    j="jiggle.exe"

    echo "Creating $j"
    $fynecmd -os windows -- -ldflags -H=windowsgui
    mv $j build/$j
  ;;
  "darwin"*)
    j="jiggle.app"

    echo "Creating $j"
    $fynecmd -os darwin
    mv $j build/$j
  ;;
  "linux"*|*"bsd"*)
    j="jiggle.tar.xz"

    echo "Creating $j"
    $fynecmd -os linux
    mv $j build/$j
    cd build
    echo "Extracting $j"
    tar -xJf $j
  ;;
  *)
    echo "Error: '$os' OS is not supported." >&2
  ;;
esac
