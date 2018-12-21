#!/usr/bin/env bash

protoc \
    -I./proto \
    --include_imports \
    --include_source_info \
    --descriptor_set_out=api_descriptor.pb \
    --go_out=plugins=grpc:proto \
    Challenge.proto
