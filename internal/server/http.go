package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	pb "infogpt/api/admin/v1"
	"infogpt/internal/conf"
	"infogpt/internal/service"
	lib "infogpt/library"

	"github.com/gin-gonic/gin"
	kgin "github.com/go-kratos/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// 包级变量，用于代理转发
var proxyHttpClient *http.Client

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, admin *service.AdminService, logger log.Logger) *khttp.Server {
	initProxyHttpClient(c)

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
	ginRouter.Any("/*path", openaiProxy)
	httpSrv.HandlePrefix("/openaiproxy", ginRouter)

	return httpSrv
}

// initProxyHttpClient 根据配置内容，初始化 proxyHttpClient
func initProxyHttpClient(httpConf *conf.Server) {
	proxyUrl := httpConf.Http.ProxyUrl
	to := httpConf.Http.ProxyTimeout.Seconds
	if to <= 0 {
		to = lib.HTTPClientProxyTimeoutS
	}

	if proxyUrl != "" {
		proxyUrl, err := url.Parse(proxyUrl)
		if err != nil {
			log.Error("parse proxy_url %s error: %v", proxyUrl, err)
			panic(err)
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		proxyHttpClient = &http.Client{
			Transport: transport,
			Timeout:   time.Duration(to) * time.Second,
		}
	} else {
		log.Info("[initProxyHttpClient] no proxy_url, use http.DefaultClient")
		proxyHttpClient = http.DefaultClient
	}
}

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
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
	newPath := strings.Replace(originPath, "/openaiproxy", "", 1)
	url = lib.OpenAIBaseAPI + newPath
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
	response, err = proxyHttpClient.Do(request)
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
