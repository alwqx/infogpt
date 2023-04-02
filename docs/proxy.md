# OpenAI API Or ChatGPT 代理

## OpenAI API

## ~~ChatGPT~~

该方案不可控，chatgpt + access_token 容易造成个人的聊天信息泄露，留了 todo，暂未实现

**access_token 和、\_puid 和浏览器绑定了**，因此配置在 InfoGPT 中的 access_token 和 puid 和在对应的浏览器上请求获取。

如果 infogpt 中的请求客户端 User-Agent 用的是 Chrome，而你的 access_token 和 puid 使用 FireFox 或者 Safari。 获取的，那么则会触发 ChatGPT 的风控。

```
Please stand by, while we are checking your browser...

Please turn JavaScript on and reload the page.
```

```json
{
  "user": {
    "id": "your user id",
    "name": "your name",
    "email": "your email",
    "image": "your image url",
    "picture": "your picture url",
    "mfa": false,
    "groups": []
  },
  "expires": "2023-05-01T09:17:58.144Z",
  "accessToken": "your_access_token"
}
```
