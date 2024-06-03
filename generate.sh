#!/usr/bin/env bash

SOURCE_PATH=../chat

if [ ! -d $SOURCE_PATH ]
then
  echo "cannot find chat path on the parent folder (${SOURCE_PATH}), do you have a copy of the API source?";
  exit 1;
fi

if ! gofumpt -version &> /dev/null
then
  echo "cannot find gofumpt in path, did you setup this repo correctly?";
  exit 1;
fi

set -ex

# cd in API repo, generate new spec and then generate code from it
( cd $SOURCE_PATH ; make openapi ; go run ./cmd/chat-manager openapi generate-client --language go-serverside --spec ./releases/v2/serverside-api.yaml --output ../stream-go )

# lint generated code with gofumpt
gofumpt -w ../stream-go