#!/usr/bin/env bash

# 执行使用 bash (sh => Bad substitution)
# 停止在运行的容器,不停止无法删除

# Dockerfile所在路径
sh_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

container=headless-shell_zh
echo checking $container

# 查询docker镜像是否存在
checkIfImageExists() {
    docker images | grep -q $container
    if [ $? -eq 0 ]; then
        echo "Image $container exists"
        return 0
    else
        echo "Image $container does not exist"
        return 1
    fi
}
searched=$(docker ps | grep $container | grep -v grep | awk '{print $2}')
if [ "$searched" != "$container" ]; then
    echo "docker ps result: " $searched
    docker stop $container
    # 没有则构建镜像,运行
    if checkIfImageExists; then
        echo "Image $container exists"
    else
        echo "Docker Image $container does not exist, now  build and run"
        docker build -t headless-shell_zh $sh_dir
    fi

    docker run -d -p 9222:9222 --rm --name $container --shm-size 2G $container

    #删除镜像 -f 强制删除
    #docker rmi -f headless-shell_zh
else
    echo "docker ps result: " $searched
fi
