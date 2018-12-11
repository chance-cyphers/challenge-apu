#!/usr/bin/env bash

protoc \
    --include_imports \
    --include_source_info \
    --proto_path=./proto \
    --descriptor_set_out=api_descriptor.pb \
    --go_out=plugins=grpc:proto \
    Challenge.proto
