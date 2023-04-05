package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	pb "infogpt/api/admin/v1"
	lib "infogpt/library"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// NewTelegramBot 根据 token 生成默认的 telegram bot
func NewTelegramBot(token string, proxyUrl string) *tgbotapi.BotAPI {
	// 判断是否使用代理http client
	var HTTPClient *http.Client
	if proxyUrl != "" {
		urlObj, err := url.Parse(proxyUrl)
		if err != nil {
			log.Panicf("parse proxy_url %s error: %v", proxyUrl, err)
		}

		transport := &http.Transport{
			Proxy: http.ProxyURL(urlObj),
		}
		HTTPClient = &http.Client{
			Transport: transport,
		}
	}

	var (
		bot *tgbotapi.BotAPI
		err error
	)

	if HTTPClient != nil {
		bot, err = tgbotapi.NewBotAPIWithClient(token, tgbotapi.APIEndpoint, HTTPClient)
	} else {
		bot, err = tgbotapi.NewBotAPI(token)
	}
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s\n", bot.Self.UserName)

	return bot
}

// AsyncProcessTelegramCommand 异步处理 telegram command
// 运行该方法后会立即返回
func (s *AdminService) AsyncProcessTelegramCommand() {
	go s.syncProcessTelegramCommand()
}

// SyncProcessTelegramCommand 同步处理telegram command
// 运行该方法后会阻塞住，异步调用请使用 AsyncProcessTelegramCommand
func (s *AdminService) syncProcessTelegramCommand() {
	u := tgbotapi.NewUpdate(0)
	// 默认请求超时60秒
	u.Timeout = 60

	updates := s.TelegramBot.GetUpdatesChan(u)
	for update := range updates {
		lib.PrintJson("DEBUG-syncProcessTelegramCommand", update)
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		fromUser := update.Message.From.UserName
		if !update.Message.IsCommand() { // ignore any non-command Messages
			s.log.Warnf("message from user %s is not command, msg=%s, skip",
				fromUser, update.Message.Text)
			continue
		}

		// 初始化返回消息结构体
		req := new(pb.SummaryReuqest)
		commandDetail := update.Message.CommandArguments()
		req.PromptDetail = commandDetail
		start := time.Now()
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /url /book and /chat"
		case "url":
			rep, err := s.UrlSummary(context.Background(), req)
			if err != nil {
				s.log.Errorf("booksummary of %s error: %v", req.PromptDetail, err)
				msg.Text = err.Error()
			} else {
				msg.Text = rep.Summary
			}
		case "book":
			rep, err := s.BookSummary(context.Background(), req)
			if err != nil {
				s.log.Errorf("booksummary of %s error: %v", req.PromptDetail, err)
				msg.Text = err.Error()
			} else {
				msg.Text = rep.Summary
			}
		case "chat":
			chatRep := &pb.OpenaiChatReuqest{
				Message: commandDetail,
			}
			rep, err := s.OpenaiChat(context.Background(), chatRep)
			if err != nil {
				s.log.Errorf("openai chat of %s error: %v", chatRep.Message, err)
				msg.Text = err.Error()
			} else {
				msg.Text = rep.Message
			}
		default:
			msg.Text = ErrNotSupportTelegramCommand.Error()
		}

		// 格式化回复信息，添加问题、耗时等
		costTime := time.Since(start)
		msg.Text = fmt.Sprintf("消息人: @%s\n消息内容: %s\n耗时: %s\n回复:\n%s",
			fromUser, commandDetail, costTime.String(), msg.Text)
		msg.ReplyToMessageID = update.Message.MessageID
		if _, err := s.TelegramBot.Send(msg); err != nil {
			s.log.Errorf("send telegram msg %v, error: %v", msg, err)
		}
	}
}

var ErrNotSupportTelegramCommand = errors.New("not support telegram bot command")
