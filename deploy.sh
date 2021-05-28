#!/usr/bin/env sh
export IMG="localhost:32000/ui-api:latest"
make docker-build docker-push IMG=${IMG}
make deploy IMG=${IMG}
