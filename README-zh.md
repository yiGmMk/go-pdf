# 支持中文的chromedp/headless-shell

    基于官方的chromedp/headless-shell,添加了中文的支持,主要通过安装字体,使用阿里云镜像优化

    ```bash
    在linux环境可以通过以下命令来运行,程序请在release页面下载或通过源码构建
    ./go-pdf -url=http://www.baidu.com -path=./baidu.pdf
    ```

## 使用

### 依赖

1. docker
2. go(for demostration)

### 使用dockerfile构建镜像

    ```bash
    bash  pdf.sh
    ```

### 拉取预构建镜像

    ```bash
    docker pull registry.cn-hangzhou.aliyuncs.com/programnotes/chromedp-headless-shell-suport-chinese:v1.2

    # run container using script
    bash run-pull-container.sh
    ```

### 测试

    ```bash
    curl -X GET http://localhost:9222/json
    ```

    通过curl调接口看到类似输出说明部署成功了
    ```bash
      [ {
         "description": "",
         "devtoolsFrontendUrl": "/devtools/inspector.html?ws=localhost:9222/devtools/page/      D2502817BA8C5F804A663645ECEA3054",
         "id": "D2502817BA8C5F804A663645ECEA3054",
         "title": "about:blank",
         "type": "page",
         "url": "about:blank",
         "webSocketDebuggerUrl": "ws://localhost:9222/devtools/page/D2502817BA8C5F804A663645ECEA3054"
      } ]
    ```
