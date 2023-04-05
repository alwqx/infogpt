package service

import (
	"context"
	"fmt"
	pb "infogpt/api/admin/v1"
	lib "infogpt/library"
	"net/url"

	"github.com/go-kratos/kratos/v2/log"
	openai "github.com/sashabaranov/go-openai"
)

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
