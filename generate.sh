#!/bin/sh
sudo rm -rf openapi
sudo docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate -i /local/types.json -g go -o /local/openapi
sudo rm openapi/go.mod openapi/go.sum
