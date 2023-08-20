#!/bin/bash

APP_NAME="messengerBot"

mkdir -p output/bin output/configs

go build -o output/bin/${APP_NAME}