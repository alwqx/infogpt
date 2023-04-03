package server

import (
	"fmt"
	"io"
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
	"github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

var adminSvc *service.AdminService

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, admin *service.AdminService, logger log.Logger) *khttp.Server {
	adminSvc = admin

	// 生成Server
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

	// 使用ratelimit
	rate, err := limiter.NewRateFromFormatted(adminSvc.AdminConf.GinRatelimitConfig)
	if err != nil {
		panic(err)
	}
	store := memory.NewStore()
	instance := limiter.New(store, rate)
	rateMiddleware := mgin.NewMiddleware(instance)

	// 使用gin框架代理 openai api
	ginRouter := gin.Default()
	ginRouter.Use(kgin.Middlewares(recovery.Recovery(), customMiddleware))
	ginRouter.Use(rateMiddleware)
	ginRouter.Any("/openaiproxy/*path", openaiProxy)
	httpSrv.HandlePrefix("/openaiproxy", ginRouter)
	ginRouter.GET("/hello", ginHello)
	httpSrv.HandlePrefix("/hello", ginRouter)
	// ChatGPT 代理相对麻烦，而且很多第三方客户端不支持，暂时 TODO
	// ginRouter.Any("/chatgptproxy/*path", chatGPTProxy)
	// httpSrv.HandlePrefix("/chatgptproxy", ginRouter)

	// syncChatGPTSession()

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
		_, err := io.Copy(w, response.Body)
		if err != nil {
			log.Errorf("io.Copy error: %v", err)
		}
		return false
	})
}

// ginHello 用于测试 gin.Router 相关功能
func ginHello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}
