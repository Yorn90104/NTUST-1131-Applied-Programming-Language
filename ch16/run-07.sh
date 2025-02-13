#!/bin/bash

GO=go
NAME=07-atomic-operations
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
