#!/usr/bin/env bash

# 执行使用 bash (sh => Bad substitution)
# 停止在运行的容器,不停止无法删除

# Dockerfile所在路径
sh_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

container=headless-shell_zh
echo checking $container

searched=$(docker ps | grep $container | grep -v grep | awk '{print $2}')
if [ "$searched" != "$container" ]; then
    echo "docker ps result: " $searched
    echo "pdf container:$container not exist, now  build and run"
    docker stop $container
    # 构建镜像,运行
    docker build -t headless-shell_zh $sh_dir
    docker run -d -p 9222:9222 --rm --name $container --shm-size 2G $container

    #删除镜像 -f 强制删除
    #docker rmi -f headless-shell_zh
else
    echo "docker ps result: " $searched
fi
