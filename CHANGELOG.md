# InfoGPT changelog

### v0.0.9

1. telegram bot 支持按照用户 id ratelimit

### v0.0.8

1. 变更配置名
2. 微信公众号根据配置信息判断是否开启
3. 拆分 servie.go 文件

### v0.0.7

1. 接入微信公众号聊天功能-比较弱，OpenAI 响应时间超过 5 秒会断开，提示"服务临时不可用"
2. telegram 回复消息加上 message ID

### v0.0.6

1. 代理 openaiproxy 接口添加 ratelimit
2. 提供 hello 接口用于 Gin 测试
3. 去掉 chatGPTProxy 相关冗余代码
4. 使用 [golangci-lint](https://github.com/golangci/golangci-lint/) 检查代码

### v0.0.5

1. ~~支持 ChatGPT、ChatGPT Plus 代理~~，不可控，chatgpt + access_token 容易造成个人的聊天信息泄露，留了 todo，暂未实现
2. 删除冗余 internal/data 目录
3. 去除多余配置，更新配置名

### v0.0.4

1. 集成 telegram bot

### v0.0.3

1. 根据网址总结内容
2. 根据书名总结内容

### v0.0.2

1. 服务器部署验证+文档

### v0.0.1

1. 初始化 api host，代理请求到 openai 服务器
2. 初始化基本聊天功能
