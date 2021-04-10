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
- 双击运行默认端口为8080
- 打开 http://localhost:8080
### (2)命令行运行
指定监听端口9999, 目录为当前目录
配置会默认加载httpserver下的`server-config.json`

```
./gohttpserver.exe -p 9999 -r ./
```

json样例
```json
{
  "port": "9999",
  "root": "webpage",
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

- 打开 http://localhost:9999
