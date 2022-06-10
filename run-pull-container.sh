#!/usr/bin/env bash

# 执行使用 bash (sh => Bad substitution)
# 停止在运行的容器,不停止无法删除

# Dockerfile所在路径
sh_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

container=registry.cn-hangzhou.aliyuncs.com/programnotes/chromedp-headless-shell-suport-chinese:v1.2
name=headless-shell_zh
echo checking $name

searched=$(docker ps | grep $container | grep -v grep | awk '{print $2}')
if [ "$searched" != "$container" ]; then
    echo "docker ps result: " $searched
    echo "pdf container:$container not exist, now  run"
    docker stop $name
    docker run -d -p 9222:9222 --rm --name $name --shm-size 2G $container

    #删除镜像 -f 强制删除
    #docker rmi -f headless-shell_zh
else
    echo "docker ps result: " $searched
fi