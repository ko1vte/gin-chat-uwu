#!/bin/bash

# 构建Docker镜像
docker build -t gin-chat-uwu:latest .

# 运行Docker容器，并进行端口映射
docker run -d -p 8000:8080 gin-chat-uwu:latest

# 输出容器ID
container_id=$(docker ps -q --filter ancestor=myapp:latest)
echo "容器已运行，ID: $container_id"