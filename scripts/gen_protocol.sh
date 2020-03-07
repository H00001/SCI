#!/usr/bin/env bash
set -e

cd ../protocol
protoc --go_out=plugins=grpc:. *.proto