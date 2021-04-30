#!/usr/bin/env bash

source "scripts/test_lib.sh"

GIT_SHA=$(git rev-parse --short HEAD || echo "GitNotFound")
if [ -n "$FAILPOINTS" ]; then
    GIT_SHA="$GIT_SHA"-FAILPOINTS
fi

VERSION_SYMBOL="${ROOT_MODULE}/api/v3/version.GitSHA"

GO_BUILD_ENV=("CGO_ENABLED=0" "GO_BUILD_FLAGS=${GO_BUILD_FLAGS}" "GOOS=${GOOS}" "GOARCH=${GOARCH}")


GO_LDFLAGS=(${GO_LDFLAGS} "-X=${VERSION_SYMBOL}=${GIT_SHA}")

echo $1