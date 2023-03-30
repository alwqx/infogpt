package service

import (
	"context"
	"net/http"
	"net/url"
	"time"

	pb "infogpt/api/admin/v1"
	"infogpt/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
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
}

func NewAdminService(openaiConf *conf.Openai, logger log.Logger) (*AdminService, error) {
	l := log.NewHelper(log.With(logger, "module", "service/admin"))

	// openai client
	openAIConfig := openai.DefaultConfig(openaiConf.ApiKey)
	if openaiConf.ProxyUrl != "" {
		proxyUrl, err := url.Parse(openaiConf.ProxyUrl)
		if err != nil {
			log.Error("parse proxy_url %s error: %v", openaiConf.ProxyUrl, err)
			return nil, err
		}
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}
		openAIConfig.HTTPClient = &http.Client{
			Transport: transport,
			Timeout:   15 * time.Second,
		}
	}
	oc := openai.NewClientWithConfig(openAIConfig)
	svc := &AdminService{
		log:          l,
		OpenAIApiKey: openaiConf.ApiKey,
		OpenAIClient: oc,
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
