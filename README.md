# Kratos Project Template For InfoGPT

## 快速部署

参考 [deploy](docs/deploy.md)

## Install Kratos

```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## Create a service

```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```

## Generate other auxiliary files by Makefile

```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```

## Automated Initialization (wire)

```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker

```bash
# build
docker build -t infogpt .

# run
docker run -d --name infogpt --rm -p 6060:6060 -p 6061:6061 -v </path/to/your/configs>:/data/conf infogpt
```

## TODOs

核心功能：

- [ ] 生成网页文章摘要
- [ ] 生成文件内容摘要
- [ ] 通过声音和练习口语

第三方机器人：

- [ ] Telegram
- [ ] Slack
- [ ] WeChat
