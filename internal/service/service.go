package service

import (
	"context"
	"net/http"
	"net/url"
	"time"

	pb "infogpt/api/admin/v1"
	"infogpt/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/wire"
	gocache "github.com/patrickmn/go-cache"
	openai "github.com/sashabaranov/go-openai"
	"github.com/silenceper/wechat/v2/officialaccount"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminService)

type AdminService struct {
	pb.UnimplementedAdminServer
	log *log.Helper

	// 暴露出来，给上层http层使用
	ProxyHttpClient *http.Client
	AdminConf       *conf.Admin

	OpenAIApiKey string
	OpenAIClient *openai.Client

	OfficialAccount    *officialaccount.OfficialAccount
	EnableWeChat       bool
	WeChatMessageCache *gocache.Cache
	WechatLimiter      *WechatLimiter

	// telegram bot 相关配置
	// 根据telegram配置项判断是否开启telegram bot功能，如果不开启，则不会运行telegram bot相关代码
	enableTelegram  bool
	TelegramBot     *tgbotapi.BotAPI
	TelegramLimiter *TelegramLimiter
}

func NewAdminService(adminConf *conf.Admin, logger log.Logger) (*AdminService, error) {
	l := log.NewHelper(log.With(logger, "module", "service/admin"))
	svc := &AdminService{
		log:                l,
		OpenAIApiKey:       adminConf.OpenaiApiKey,
		AdminConf:          adminConf,
		WeChatMessageCache: gocache.New(2*time.Hour, time.Hour), // 过期时间2小时，检查周期1小时
	}

	// openai client
	openAIConfig := openai.DefaultConfig(adminConf.OpenaiApiKey)
	if adminConf.ProxyUrl != "" {
		proxyUrl, err := url.Parse(adminConf.ProxyUrl)
		if err != nil {
			log.Errorf("parse proxy_url %s error: %v", adminConf.ProxyUrl, err)
			return nil, err
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		svc.ProxyHttpClient = &http.Client{
			Transport: transport,
		}
		openAIConfig.HTTPClient = svc.ProxyHttpClient
	} else {
		l.Info("[NewAdminService] no proxy_url, use http.DefaultClient")
		svc.ProxyHttpClient = http.DefaultClient
	}
	svc.OpenAIClient = openai.NewClientWithConfig(openAIConfig)

	// 公众号
	svc.OfficialAccount = NewOfficialAccount(adminConf)
	svc.EnableWeChat = svc.isOfficialAccountEnable(adminConf.Wechat)
	svc.WechatLimiter = NewWechatLimiter(adminConf.Wechat)

	// 开始异步处理 telegram command
	svc.enableTelegram = (adminConf.Telegram.Token != "")
	if svc.enableTelegram {
		svc.TelegramBot = NewTelegramBot(adminConf.Telegram.Token, adminConf.ProxyUrl)
		svc.TelegramLimiter = NewTelegramLimiter(adminConf.Telegram)
		svc.AsyncProcessTelegramCommand()
	} else {
		log.Warn("[NewAdminService] not enable telegram bot, skip")
	}

	return svc, nil
}

func (s *AdminService) HealthCheck(_ context.Context, req *pb.HealthRequest) (*pb.HealthReply, error) {
	resp := &pb.HealthReply{
		Message: "ok",
	}
	return resp, nil
}

func (s *AdminService) AppInfo(_ context.Context, req *pb.AppInfoRequest) (*pb.AppInfoReply, error) {
	resp := &pb.AppInfoReply{
		Version: "v0.0.1",
	}
	return resp, nil
}
