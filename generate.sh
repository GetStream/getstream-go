#!/usr/bin/env bash

DST_PATH=`pwd`
SOURCE_PATH=../chat

if [ ! -d $SOURCE_PATH ]
then
  echo "cannot find chat path on the parent folder (${SOURCE_PATH}), do you have a copy of the API source?";
  exit 1;
fi

if ! command -v jq >/dev/null 2>&1
then
  echo "jq is required to resolve the latest protocol release.";
  exit 1;
fi

set -ex

# The spec comes from the latest protocol release rather than from chat's working tree, so the
# generated code always corresponds to a tag someone else can regenerate from. protocol carries
# other tag namespaces, hence the openapi- prefix filter. protocol is public, so no auth is needed.
# The generator binary still comes from chat master, which is why .spec-version records it too.
SPEC_TAG=$(curl -fsSL -H "Accept: application/vnd.github+json" https://api.github.com/repos/GetStream/protocol/releases \
  | jq -r 'map(select(.tag_name | startswith("openapi-"))) | .[0].tag_name')
[ -n "$SPEC_TAG" ] && [ "$SPEC_TAG" != "null" ] || { echo "could not resolve the latest protocol openapi release"; exit 1; }
SPEC_DIR=$(mktemp -d)
trap 'rm -rf "$SPEC_DIR"' EXIT
SPEC_FILE="$SPEC_DIR/serverside-api.yaml"
curl -fsSL -o "$SPEC_FILE" "https://raw.githubusercontent.com/GetStream/protocol/${SPEC_TAG}/openapi/v2/serverside-api.yaml"

# cd in API repo, build the generator and then generate code from the release spec
( cd $SOURCE_PATH ; make -C projects/chat-manager build ; ./build/chat-manager openapi generate-client --language go-serverside --spec "$SPEC_FILE" --output $DST_PATH ; ./build/chat-manager openapi generate-webhook-fixtures --output $DST_PATH/tests/fixtures/webhooks --time-format=unix-ns )

printf 'spec: %s\ngenerator: chat %s\n' "$SPEC_TAG" "$(git -C $SOURCE_PATH describe --tags --always)" > "$DST_PATH/.spec-version"

./lint.sh
