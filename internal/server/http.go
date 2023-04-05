package server

import (
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
	"github.com/silenceper/wechat/v2/officialaccount/message"
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

	// 设置路由
	ginRouter.GET("/hello", ginHello)
	ginRouter.Any("/openaiproxy/*path", openaiProxy)
	ginRouter.Any("/officialaccount", processOfficialAccountMessage)
	httpSrv.HandlePrefix("/hello", ginRouter)
	httpSrv.HandlePrefix("/openaiproxy", ginRouter)
	httpSrv.HandlePrefix("/officialaccount", ginRouter)

	return httpSrv
}

// openaiProxy 将请求转发到openai api 服务器地址
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

// ServeOfficialAccountMessage 处理微信公众号消息
func processOfficialAccountMessage(ctx *gin.Context) {
	server := adminSvc.OfficialAccount.GetServer(ctx.Request, ctx.Writer)
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		// 先判断是否是关键字自动回复
		respText := new(message.Text)
		if replyInfo, ok := adminSvc.AdminConf.OfficialAccount.AutoReplay[msg.Content]; ok {
			respText.Content = message.CDATA(replyInfo)
		} else {
			chatReq := &pb.OpenaiChatReuqest{
				Message: msg.Content,
			}
			// 如果消息过于复杂，OpenAI处理时间超过5秒，微信会断开连接，并且重试3次
			// 这里要做好检查，如果超过5秒，就返回success或者提示信息
			// 然后把问题缓存起来，新起一个逻辑，去请求答案
			// 详情见: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Passive_user_reply_message.html
			chatResp, err := adminSvc.OpenaiChat(ctx, chatReq)
			if err != nil {
				log.Errorf("[processOfficialAccountMessage] serveOfficialWechat error: %v", err)
				respText.Content = message.CDATA(err.Error())
			} else {
				respText.Content = message.CDATA(chatResp.Message)
			}
		}

		return &message.Reply{MsgType: message.MsgTypeText, MsgData: respText}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		log.Errorf("[processOfficialAccountMessage] serve.Serve error: %v", err)
		return
	}
	err = server.Send()
	if err != nil {
		log.Errorf("[processOfficialAccountMessage] server.Send error: %v", err)
	}
}
