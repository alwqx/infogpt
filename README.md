# InfoGPT - Made ChatGPT/LLM easy

## 对于个人用户

1. OpenAI API 代理
2. ~~共享 API Key~~

## Telegram Bot

## WeChat 公众号

## 部署

得益于 Golang 语言的静态编译，部署运行只需 3 步：

1. 编译/下载得到二进制文件
2. 得到正确的配置
3. 配置域名-可选择

详情参考文档 [deploy](docs/deploy.md)

## Docker 快速部署

```bash
# build
docker build -t infogpt .

# run
docker run -d --name infogpt --rm -p 6060:6060 -p 6061:6061 -v </path/to/your/configs>:/data/conf infogpt
```

## TODOs

- [x] 支持 OpenAI Proxy
- [x] REST/gRPC 接口
  - [x] 聊天
  - [x] 生成网页文章摘要
  - [x] 生成书籍内容摘要
- [x] Telegram
  - [x] /chat /url /book 三个命令
  - [x] 使用频率限制
- [ ] WeChat 公众号有限支持
  - [ ] 使用频率限制
  - [ ] 超时缓存问答
- [ ] 支持 ChatGPT Proxy
- [ ] Slack
- [ ] 生成文件内容摘要
