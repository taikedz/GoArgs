#!/usr/bin/env sh
podman run --rm -it -v "$PWD:/hostdata" docker.io/library/golang:1.24-alpine sh -c "cd /hostdata; $*"
