#!/bin/bash

TARGET_DIR="generate"

GENERATE_DIR="./cmd/$TARGET_DIR"

cd "$GENERATE_DIR" || exit

echo "Start Generating"

go run . -dsn 'root:123456@tcp(localhost:13306)/messengerBot?charset=utf8mb4&parseTime=True'