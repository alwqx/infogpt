# InfoGPT - Made ChatGPT/LLM easy

<p align="center">
  <a href="https://pkg.go.dev/badge/github.com/alwqx/infogpt" title="GoDoc">
    <img src="https://pkg.go.dev/badge/github.com/alwqx/infogpt?status.svg">
  </a>
  <a href="https://github.com/alwqx/infogpt/releases" title="GitHub release">
    <img src="https://github.com/alwqx/infogpt/releases.svg">
  </a>
  <a href="https://opensource.org/licenses/MIT" title="License: MIT">
    <img src="https://img.shields.io/badge/License-MIT-blue.svg">
  </a>
  <a href="https://www.tickgit.com/browse?repo=github.com/alwqx/infogpt&branch=main" title="TODOs">
    <img src="https://badgen.net/https/api.tickgit.com/badgen/github.com/alwqx/infogpt/main">
  </a>
</p>

## 对于个人用户

1. OpenAI API 代理
2. ~~共享 API Key~~

![](https://alwq.site/github/infogpt_openai_proxy.gif)

## Telegram Bot

![](https://alwq.site/github/infogpt_telegram.gif)

## WeChat 公众号

TODOs

## Docker 快速部署

```bash
# build
docker build -t infogpt .

# run
docker run -d --name infogpt --rm -p 6060:6060 -p 6061:6061 -v </path/to/your/configs>:/data/conf infogpt:v0.0.12
```

## 部署

得益于 Golang 语言的静态编译，部署运行只需 3 步：

1. 编译/下载得到二进制文件
2. 得到正确的配置
3. 配置域名-可选择

详情参考文档 [deploy](docs/deploy.md)

## TODOs

- [x] 支持 OpenAI Proxy
- [x] REST/gRPC 接口
  - [x] 聊天
  - [x] 生成网页文章摘要
  - [x] 生成书籍内容摘要
- [x] Telegram
  - [x] /chat /url /book 三个命令
  - [x] 使用频率限制
- [x] WeChat 公众号有限支持
  - [x] 使用频率限制
  - [x] 超时缓存问答
- [ ] 支持 ChatGPT Proxy
- [ ] Slack
- [ ] 生成文件内容摘要

## 致谢

本开源项目的开发，离不开这些开源项目的支持：

- [go-kratos](https://github.com/go-kratos/kratos)
