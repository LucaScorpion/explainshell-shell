#!/usr/bin/env bash
set -eu

GOARCH=amd64
VERSION=$(git describe --tags)
APP="./cmd/explainshell.go"
CGO_ENABLED=0

echo 'Building explainshell...'

for GOOS in linux darwin windows
do
  echo "- $GOOS"

  OUTPUT="bin/explainshell-$VERSION-$GOOS-$GOARCH"
  if [ $GOOS = 'windows' ]
  then
      OUTPUT="$OUTPUT.exe"
  fi

  go build -o "$OUTPUT" "$APP"
done

echo 'Done!'
