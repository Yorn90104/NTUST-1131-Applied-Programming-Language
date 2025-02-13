#!/bin/bash

GO=go
NAME=06-race-conditions
SRC=${NAME}.go
BIN=./${NAME}
OUTPUT=3

${GO} build ${SRC}

while true; do
  res=`${BIN}`
  echo $res
  if [[ "${res}" != "${OUTPUT}" ]]; then
    break
  fi
done
