package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	lib "infogpt/library"

	tls_client "github.com/bogdanfinn/tls-client"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// 包级变量，用于代理转发
var (
	// 个人在浏览器请求 https://chat.openai.com/api/auth/session 返回的结果中的值
	chatGPTAccessToken string
	chatGPTPuid        string
	jar                = tls_client.NewCookieJar()
	puid               string
)

func customMiddleware(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
		if tr, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("operation:", tr.Operation())
		}
		reply, err = handler(ctx, req)
		return
	}
}

// syncChatGPTSession 同步 ChatGPT 会话信息，用于代理
func syncChatGPTSession() {
	go func() {
		path := lib.OpenAIChatGPTAPI + "/models"
		req, err := http.NewRequest(http.MethodGet, path, nil)
		if err != nil {
			log.Errorf("NewRequest with path %s error: %v", err)
		}
		setChatGPTReqHeaders(req)
		// 初始化puid，后面使用定期请求中的
		puid = adminSvc.AdminConf.ChatgptPuid
		// Initial puid cookie
		req.AddCookie(
			&http.Cookie{
				Name:  "_puid",
				Value: puid,
			},
		)
		for {
			resp, err := adminSvc.ProxyHttpClient.Do(req)
			if err != nil {
				break
			}
			defer resp.Body.Close()
			if resp.StatusCode != 200 {
				body, _ := io.ReadAll(resp.Body)
				log.Errorf("request error, status_code=%d, body=%s", resp.StatusCode, string(body))
				break
			}
			cookies := resp.Cookies()
			// Find _puid cookie
			for _, cookie := range cookies {
				if cookie.Name == "_puid" {
					puid = cookie.Value
					break
				}
			}
			time.Sleep(6 * time.Hour)
		}
		log.Error("request req error: %v", err)
	}()
}

func setChatGPTReqHeaders(req *http.Request) {
	if req == nil {
		log.Error("setChatGPTReqHeaders req is nil, skip")
		return
	}

	req.Header.Set("Host", "chat.openai.com")
	req.Header.Set("origin", "https://chat.openai.com/chat")
	req.Header.Set("referer", "https://chat.openai.com/chat")
	// req.Header.Set("sec-ch-ua", `Chromium";v="110", "Not A(Brand";v="24", "Brave";v="110`)
	// req.Header.Set("sec-ch-ua-platform", "Linux")
	req.Header.Set("content-type", "application/json")
	req.Header.Set("accept", "text/event-stream")
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/111.0")
	// Set authorization header
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", adminSvc.AdminConf.ChatgptAccessToken))
}
