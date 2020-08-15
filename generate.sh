#!/bin/sh
rm -rf openapi
docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/types.json -g go -o /local/openapi
rm openapi/go.mod openapi/go.sum
