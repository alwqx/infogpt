# InfoGPT changelog

### v0.0.14

1. 更新文档，部署信息更详细
2. 调整 release action，把配置目录`configs`整体打包，而非打包单个文件

### v0.0.13

1. 修复 release action 配置错误

### v0.0.12

1. 更新 Readme，新增 telegram 和 openai proxy gif
2. 新增 github go release action
3. 新增 docker image infogpt:v0.0.12

### v0.0.11

1. 公众号 ratelimit
   oai.infogpt.cc/openaiproxy

### v0.0.10

1. 没有开启企业认证的微信公众号，聊天消息超时缓存
   说明：根据微信文档 [被动回复用户消息](https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Passive_user_reply_message.html)，第三方服务要在 5 秒内响应用户的消息，而聊天模型的响应时间**普便**超过 5 秒，因为为了`提供较好的体验`，超过 5 秒没有响应，我们会把用户的聊天缓存，等有结果了再返回
2. 返回消息过长，自动分割

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
