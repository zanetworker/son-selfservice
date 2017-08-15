#!/usr/bin/env bash

IMAGE="selfservice-backend"

docker build --rm -f Dockerfile.compile . --tag $IMAGE

# remove binary if exists
FILE="selfservice-backend"
if [ -e "$FILE" ];
then
  rm -f "$FILE"
fi

docker run --rm -it -v "$GOPATH":/gopath -e "FILE=$FILE" -v "$(pwd)":/$FILE \
  -w "/$FILE" $IMAGE sh -c 'go build -a --installsuffix cgo --ldflags="-s" -o $FILE'
