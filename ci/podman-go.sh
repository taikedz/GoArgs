#!/usr/bin/env sh

# Run unit tests:
#    bash podman-go.sh
#
# Run any other actions:
#    bash podman-go.sh COMMAND ...

HERE="$(dirname "$0")" # <repo>/ci/
ROOT="$(readlink -f "$HERE/..")" # /.../<repo>/

if [[ -z "$*" ]]; then
    actions="cd unittests; go test"
else
    actions="$*"
fi

podman run --rm -it -v "$ROOT:/hostdata" docker.io/library/golang:1.24-alpine sh -c "cd /hostdata; $actions"
