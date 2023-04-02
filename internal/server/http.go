package server

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	pb "infogpt/api/admin/v1"
	"infogpt/internal/conf"
	"infogpt/internal/service"
	lib "infogpt/library"

	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

var adminSvc *service.AdminService

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, admin *service.AdminService, logger log.Logger) *khttp.Server {
	adminSvc = admin

	var opts = []khttp.ServerOption{
		khttp.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, khttp.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, khttp.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, khttp.Timeout(c.Http.Timeout.AsDuration()))
	}
	httpSrv := khttp.NewServer(opts...)
	// 设置swagger api
	swaggerHandler := openapiv2.NewHandler()
	httpSrv.HandlePrefix("/q", swaggerHandler)

	// 注册admin服务
	pb.RegisterAdminHTTPServer(httpSrv, admin)

	// 使用gin框架代理 openai api
	ginRouter := gin.Default()
	ginRouter.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))
	ginRouter.Any("/openaiproxy/*path", openaiProxy)
	httpSrv.HandlePrefix("/openaiproxy", ginRouter)
	// ChatGPT 代理相对麻烦，而且很多第三方客户端不支持，暂时 TODO
	// ginRouter.Any("/chatgptproxy/*path", chatGPTProxy)
	// httpSrv.HandlePrefix("/chatgptproxy", ginRouter)

	syncChatGPTSession()

	return httpSrv
}

// openaiProxy 将请求转发到openai api 服务器地址
// refer https://github.com/acheong08/ChatGPT-Proxy-V4/blob/main/main.go
func openaiProxy(ctx *gin.Context) {
	var (
		url           string
		err           error
		requestMethod string
		request       *http.Request
		response      *http.Response
	)

	// 把代理的域名前缀 /openaiproxy 去掉
	originPath := ctx.Param("path")
	// newPath := strings.Replace(originPath, "/", "", 1)
	newPath := originPath
	url = lib.OpenAIBaseAPI + newPath
	fmt.Printf("DEBUG chatGPTProxy originPath=%s, newPath=%s\n, url=%s", originPath, newPath, url)
	requestMethod = ctx.Request.Method

	request, err = http.NewRequest(requestMethod, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for key, values := range ctx.Request.Header {
		// 指定 Accept-Encoding 不可控，会导致openai返回的信息乱码
		// 比如 Accept-Encoding： "gzip, deflate, br"，返回就乱码
		if key == "Accept-Encoding" {
			continue
		}
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}

	// 发送请求
	response, err = adminSvc.ProxyHttpClient.Do(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()
	ctx.Header("Content-Type", response.Header.Get("Content-Type"))
	ctx.Status(response.StatusCode)
	ctx.Stream(func(w io.Writer) bool {
		// Write data to client
		io.Copy(w, response.Body)
		return false
	})
}

// chatGPTProxy 将 ChatGPT、ChatGPT Plus请求代理到 https://chat.openai.com/backend-api
// ChatGPT 代理相对麻烦，而且很多第三方客户端不支持，暂时 TODO
func chatGPTProxy(ctx *gin.Context) {
	var (
		url           string
		err           error
		requestMethod string
		request       *http.Request
		response      *http.Response
	)

	// 把代理的域名前缀 /chatgptproxy 去掉
	originPath := ctx.Param("path")
	newPath := originPath
	// newPath := strings.Replace(originPath, "/chatgptproxy", "", 1)
	url = lib.OpenAIChatGPTAPI + newPath
	fmt.Printf("DEBUG chatGPTProxy originPath=%s, newPath=%s, url=%s\n", originPath, newPath, url)
	requestMethod = ctx.Request.Method

	lib.PrintJson("DEBUG-chatGPTProxy-originRequest.Header", ctx.Request.Header)
	lib.PrintJson("DEBUG-chatGPTProxy-originRequest.URL", ctx.Request.URL)

	request, err = http.NewRequest(requestMethod, url, ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	for key, values := range ctx.Request.Header {
		// 指定 Accept-Encoding 不可控，会导致openai返回的信息乱码
		// 比如 Accept-Encoding： "gzip, deflate, br"，返回就乱码
		if key == "Accept-Encoding" {
			continue
		}
		for _, v := range values {
			request.Header.Add(key, v)
		}
	}

	request.Header.Set("Host", "chat.openai.com")
	request.Header.Set("Origin", "https://chat.openai.com/chat")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Connection", "keep-alive")
	request.Header.Set("Keep-Alive", "timeout=360")
	request.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	authParam := ctx.Request.Header.Get("Authorization")
	if authParam != "" {
		request.Header.Set("Authorization", authParam)
	} else {
		log.Warn("[chatGPTProxy] request not provide Authorization header, use admin cofig OpenAIAccessToken")
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", adminSvc.AdminConf.ChatgptAccessToken))
	}

	if ctx.Request.Header.Get("Puid") == "" {
		request.AddCookie(
			&http.Cookie{
				Name:  "_puid",
				Value: puid,
			},
		)
	} else {
		request.AddCookie(
			&http.Cookie{
				Name:  "_puid",
				Value: ctx.Request.Header.Get("Puid"),
			},
		)
	}

	// 打印 req
	lib.PrintJson("DEBUG-chatGPTProxy-proxyRequest.Header", request.Header)
	lib.PrintJson("DEBUG-chatGPTProxy-proxyRequest.URL", request.URL)
	// 发送请求
	response, err = adminSvc.ProxyHttpClient.Do(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer response.Body.Close()

	// 检查响应码
	// 取 request body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if response.StatusCode != http.StatusOK {
		log.Warnf("[chatGPTProxy] proxy response code=%d, body=%s\n", http.StatusOK, string(body))
	}

	response.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	ctx.Header("Content-Type", response.Header.Get("Content-Type"))
	ctx.Status(response.StatusCode)
	ctx.Stream(func(w io.Writer) bool {
		// Write data to client
		io.Copy(w, response.Body)
		return false
	})
}
