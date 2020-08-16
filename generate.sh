#!/bin/sh
rm -rf client
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/types.json -g go -o /local/client --package-name=client
rm client/go.mod client/go.sum
