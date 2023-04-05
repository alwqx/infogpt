package service

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	pb "infogpt/api/admin/v1"
	"infogpt/internal/conf"
	lib "infogpt/library"

	"github.com/go-kratos/kratos/v2/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/google/wire"
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

	OfficialAccount *officialaccount.OfficialAccount

	// telegram bot 相关配置
	// 根据telegram配置项判断是否开启telegram bot功能，如果不开启，则不会运行telegram bot相关代码
	enableTelegram bool
	TelegramBot    *tgbotapi.BotAPI
}

func NewAdminService(adminConf *conf.Admin, logger log.Logger) (*AdminService, error) {
	for k, v := range adminConf.OfficialAccount.AutoReplay {
		fmt.Printf("DEBUG key=%s, v=%s", k, v)
	}
	l := log.NewHelper(log.With(logger, "module", "service/admin"))
	svc := &AdminService{
		log:          l,
		OpenAIApiKey: adminConf.OpenaiApiKey,
		AdminConf:    adminConf,
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

	// 开始异步处理 telegram command
	svc.enableTelegram = (adminConf.TelegramToken != "")
	if svc.enableTelegram {
		svc.TelegramBot = NewTelegramBot(adminConf.TelegramToken, adminConf.ProxyUrl)
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

func (s *AdminService) OpenaiChat(ctx context.Context, req *pb.OpenaiChatReuqest) (*pb.OpenaiChatReply, error) {
	chatReq := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: req.Message,
			},
		},
	}
	resp, err := s.OpenAIClient.CreateChatCompletion(ctx, chatReq)
	if err != nil {
		log.Errorf("chat with openai error: %v", err)
		return nil, err
	}

	reply := &pb.OpenaiChatReply{
		Message: resp.Choices[0].Message.Content,
	}
	return reply, nil
}

// UrlSummary 抓取 url 内容，调用OpenAI的模型生成内容摘要
func (s *AdminService) UrlSummary(ctx context.Context, req *pb.SummaryReuqest) (*pb.SummaryReply, error) {
	// 检查url
	_, err := url.Parse(req.PromptDetail)
	if err != nil {
		return nil, err
	}

	chatCnt := fmt.Sprintf("%s %s", lib.UrlSummaryPromptCN, req.PromptDetail)
	log.Infof("[AdminService][UrlSummary] prompt is %s", chatCnt)
	chatReq := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: chatCnt,
			},
		},
	}

	resp, err := s.OpenAIClient.CreateChatCompletion(ctx, chatReq)
	if err != nil {
		log.Errorf("chat with openai error: %v", err)
		return nil, err
	}
	reply := &pb.SummaryReply{
		Summary: resp.Choices[0].Message.Content,
	}

	return reply, nil
}

// BookSummary 根据书名，调用OpenAI的模型生成内容摘要
func (s *AdminService) BookSummary(ctx context.Context, req *pb.SummaryReuqest) (*pb.SummaryReply, error) {
	chatCnt := fmt.Sprintf("%s %s", lib.BookSummaryPromptCN, req.PromptDetail)
	log.Infof("[AdminService][BookSummary] prompt is %s", chatCnt)
	chatReq := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: chatCnt,
			},
		},
	}

	resp, err := s.OpenAIClient.CreateChatCompletion(ctx, chatReq)
	if err != nil {
		log.Errorf("chat with openai error: %v", err)
		return nil, err
	}
	reply := &pb.SummaryReply{
		Summary: resp.Choices[0].Message.Content,
	}

	return reply, nil
}
