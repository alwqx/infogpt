# 部署

## 服务器部署 docker

1. 构建镜像

```bash
# build
docker build -t infogpt .

# run
docker run -d --name infogpt --rm -p 6060:6060 -p 6061:6061 -v </path/to/your/configs>:/data/conf infogpt
```

2. 更新配置为自己的，保存并记住路径

```bash
server:
  http:
    addr: 0.0.0.0:6060
    timeout: 15s
    proxy_url: ""
    proxy_timeout: 15s
  grpc:
    addr: 0.0.0.0:6061
    timeout: 2s
openai:
  api_key: ""
  proxy_url: ""
data:
  database:
    driver: mysql
    source: root:root@tcp(127.0.0.1:3306)/test
  redis:
    addr: 127.0.0.1:6379
    read_timeout: 0.2s
    write_timeout: 0.2s
```

3. 运行容器

```bash
# run
docker run -d --name infogpt --rm -p 6060:6060 -p 6061:6061 -v configs:/data/conf infogpt
```

4. 查看日志

```bash
$ docker logs -f infogpt
2023/03/31 03:21:13 maxprocs: Leaving GOMAXPROCS=2: CPU quota undefined
DEBUG msg=config loaded: config.yaml format: yaml
INFO msg=[initProxyHttpClient] no proxy_url, use http.DefaultClient
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /*path                    --> infogpt/internal/server.openaiProxy (4 handlers)
[GIN-debug] POST   /*path                    --> infogpt/internal/server.openaiProxy (4 handlers)
[GIN-debug] PUT    /*path                    --> infogpt/internal/server.openaiProxy (4 handlers)
[GIN-debug] PATCH  /*path                    --> infogpt/internal/server.openaiProxy (4 handlers)
[GIN-debug] HEAD   /*path                    --> infogpt/internal/server.openaiProxy (4 handlers)
[GIN-debug] OPTIONS /*path                    --> infogpt/internal/server.openaiProxy (4 handlers)
[GIN-debug] DELETE /*path                    --> infogpt/internal/server.openaiProxy (4 handlers)
[GIN-debug] CONNECT /*path                    --> infogpt/internal/server.openaiProxy (4 handlers)
[GIN-debug] TRACE  /*path                    --> infogpt/internal/server.openaiProxy (4 handlers)
INFO ts=2023-03-31T03:21:13Z caller=server.go:276 service.id=75f619a25e4d service.name= service.version=1e2e643 trace.id= span.id= msg=[HTTP] server listening on: [::]:6060
INFO ts=2023-03-31T03:21:13Z caller=server.go:193 service.id=75f619a25e4d service.name= service.version=1e2e643 trace.id= span.id= msg=[gRPC] server listening on: [::]:6061
```
