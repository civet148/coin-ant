#!/bin/bash

# MySQL数据源(正式环境需修改成实际数据库配置)
DSN="mysql://root:123456@172.17.0.1:3306/coin-ant?charset=utf8mb4"

IMAGE_URL='coin-ant:latest'
CONTAINER_NAME=coin-ant

docker rm -f $CONTAINER_NAME
docker run --restart always --name $CONTAINER_NAME -d $IMAGE_URL coin-ant run --dsn "${DSN}" 
