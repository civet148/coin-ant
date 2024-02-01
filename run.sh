#!/bin/bash

# MySQL数据源(正式环境需修改成实际数据库配置)
DSN="mysql://root:123456@192.168.1.16:3306/coin-ant?charset=utf8mb4"

# 集群管理系统HTTP服务监听地址
LISTEN_ADDR="0.0.0.0:8008"

./coin-ant run --debug --dsn "${DSN}"  $LISTEN_ADDR

