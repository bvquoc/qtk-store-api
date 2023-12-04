#!/bin/bash

echo "Building application"
go build
kill -15 $(pidof qtk-store-api)
#nohup ./qtk-store-api > /dev/null 2>&1&
pm2 stop 0
pm2 start qtk-store-api
echo "Deploy completed"