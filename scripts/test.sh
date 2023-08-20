#!/bin/bash

export GIN_MODE=debug

mkdir -p logs
GO_OUT=./logs/start.log  #应用的启动日志

echo "starting run test app"
nohup ./output/messengerBot -env "test" > ${GO_OUT} 2>&1 &
echo "started"