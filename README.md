# chromedp-headless-shell-suport-chinese

using docker to deploy chromedp/headless-shell,and add suport for chinese(install fonts,using mirror etc.)

## how to use

### depedency

1. docker
2. go(for demostration)

### build docker image and using and run

    ```bash
    bash  pdf.sh
    ```

### pull pre-build image and run

    ```bash
    docker pull registry.cn-hangzhou.aliyuncs.com/programnotes/chromedp-headless-shell-suport-chinese:v1.2

    docker run -d registry.cn-hangzhou.aliyuncs.com/programnotes/chromedp-headless-shell-suport-chinese:v1.2
    ```

### test

    ```bash
    curl -X GET http://localhost:9222/json
    ```

    if you see the result,it means the server is running.
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
