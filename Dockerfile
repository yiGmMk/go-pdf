FROM chromedp/headless-shell


### 打印中文依赖中文字体 ttf-wqy-zenhei fonts-wqy-microhei
### 字体管理包依赖  fontconfig xfonts-utils
# 参考/references
# 1.https://www.cnblogs.com/Jimc/p/10302267.html
# 2.https://www.cnblogs.com/igoodful/p/11235004.html
# 3.镜像源: https://developer.aliyun.com/article/765348
# 4.chromedp: https://mojotv.cn/go/chromedp-example

RUN rm /etc/apt/sources.list
RUN echo 'deb http://mirrors.aliyun.com/debian/ buster main non-free contrib\n' >> /etc/apt/sources.list \
    'deb-src http://mirrors.aliyun.com/debian/ buster main non-free contrib\n' >> /etc/apt/sources.list \
    'deb http://mirrors.aliyun.com/debian-security buster/updates main\n' >> /etc/apt/sources.list \
    'deb-src http://mirrors.aliyun.com/debian-security buster/updates main\n' >> /etc/apt/sources.list \
    'deb http://mirrors.aliyun.com/debian/ buster-updates main non-free contrib\n' >> /etc/apt/sources.list \
    'deb-src http://mirrors.aliyun.com/debian/ buster-updates main non-free contrib\n' >> /etc/apt/sources.list \
    'deb http://mirrors.aliyun.com/debian/ buster-backports main non-free contrib\n' >> /etc/apt/sources.list \
    'deb-src http://mirrors.aliyun.com/debian/ buster-backports main non-free contrib\n' >> /etc/apt/sources.list

# 安装中文字体
RUN apt-get update -y \
    && apt install -y apt-transport-https ca-certificates \
    && apt-get install -y fonts-noto fonts-noto-cjk ttf-wqy-zenhei fonts-wqy-microhei fontconfig xfonts-utils

# 删除apt缓存 clean apt cache
RUN apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/