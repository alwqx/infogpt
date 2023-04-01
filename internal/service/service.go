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
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminService)

type AdminService struct {
	pb.UnimplementedAdminServer

	log *log.Helper

	OpenAIApiKey string
	OpenAIClient *openai.Client

	TelegramBot *tgbotapi.BotAPI
}

func NewAdminService(adminConf *conf.Admin, logger log.Logger) (*AdminService, error) {
	l := log.NewHelper(log.With(logger, "module", "service/admin"))

	// openai client
	openAIConfig := openai.DefaultConfig(adminConf.ApiKey)
	if adminConf.ProxyUrl != "" {
		proxyUrl, err := url.Parse(adminConf.ProxyUrl)
		if err != nil {
			log.Error("parse proxy_url %s error: %v", adminConf.ProxyUrl, err)
			return nil, err
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		openAIConfig.HTTPClient = &http.Client{
			Transport: transport,
		}
	}
	oc := openai.NewClientWithConfig(openAIConfig)

	svc := &AdminService{
		log:          l,
		OpenAIApiKey: adminConf.ApiKey,
		OpenAIClient: oc,
		TelegramBot:  NewTelegramBot(adminConf.TelegramToken, adminConf.ProxyUrl),
	}

	// 开始异步处理 telegram command
	svc.AsyncProcessTelegramCommand()

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
