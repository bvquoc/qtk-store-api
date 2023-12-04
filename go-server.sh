#!/bin/bash
mkdir -p ./storage
chmod 777 -R ./storage

# Set the GOFLAGS environment variable to include -buildvcs=false
export GOFLAGS="-buildvcs=false"

# Run gin
gin --appPort 8080 --immediate
