#!/usr/bin/env bash

set -euxo pipefail

GOBIN=$(pwd)/bin go insatll \
         github.com/gunk/gunk

$(pwd)/bin/gunk generate ./...