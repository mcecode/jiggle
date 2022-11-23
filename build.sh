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
    echo "Creating jiggle.exe"
    $fynecmd -os windows -- -ldflags -H=windowsgui
    mv jiggle.exe build/jiggle.exe
  ;;
  "darwin"*)
    echo "Creating jiggle.app"
    $fynecmd -os darwin
    mv jiggle.app build/jiggle.app
  ;;
  "linux"*|*"bsd"*)
    echo "Creating tmp-pkg"
    $fynecmd -os linux
    mv tmp-pkg build/tmp-pkg
  ;;
  *)
    echo "Error: '$os' OS is not supported." >&2
  ;;
esac
