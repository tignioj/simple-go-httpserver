# 基于golang编写的简易HTTP服务器

# 用法
## 1. 下载项目到本地
```
git clone https://github.com/tignioj/simple-go-httpserver.git
```

## 2. 编译
```
go build
```

## 3. 运行
### (1)双击运行
- 双击运行默认端口为8088
- 打开 http://localhost:8088
### (2)命令行运行
指定监听端口9999, 目录为当前目录
配置会默认加载httpserver下的`server-config.json`

```
./gohttpserver.exe -p 9999 -r ./
```

默认配置文件如下
root 为网页指定目录
server-config.json
```json
{
  "port": "8088",
  "root": "./webpage",
  "content_type": {
    "html": "text/html",
    "css": "text/css",
    "woff": "font/woff2",
    "js": "text/javascript",
    "svg": "image/svg+xml",
    "ico": "image/x-icon"
  },
  "header": {
    "user": "zs",
    "pwd": "123",
    "secret": "sss"
  }
}
```

你也可以指定配置文件
```
./gohttpserver.exe -c myconfig.json
```
查看帮助 -h
该帮助文档由 https://github.com/tignioj/go-get-argsmap-from-commandline 生成, 帮助文档的配置文件为`help.json`
```

Using config file: server-config.json
Usage:
|------------|------------------------------|-------------|--------------------
| flag       | usage                        | expect      | default
|------------|------------------------------|-------------|--------------------
| -h         | show help                    |             |
| -p         | server port                  | pure number | 8080
| -r         | web root                     |             | ./
| -c         | path to server configuration |             | server-config.json
|------------|------------------------------|-------------|--------------------

```