package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	pb "infogpt/api/admin/v1"
	"infogpt/internal/conf"
	lib "infogpt/library"

	"github.com/go-kratos/kratos/v2/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// TelegramLimiter 控制 telegram rate 的limiter
type TelegramLimiter struct {
	ExcludeKeyMap map[string]struct{}
	Limiter       *limiter.Limiter
	UserLimiter   *limiter.Limiter
}

func NewTelegramLimiter(telegramConf *conf.Telegram) *TelegramLimiter {
	tl := new(TelegramLimiter)
	var store limiter.Store

	// 判断是否开启限流并赋值
	if telegramConf.Ratelimit != "" {
		store = memory.NewStore()
		rate, err := limiter.NewRateFromFormatted(telegramConf.Ratelimit)
		if err != nil {
			panic(err)
		}
		tl.Limiter = limiter.New(store, rate)
	}

	if telegramConf.UserRatelimit != "" {
		if store == nil {
			store = memory.NewStore()
		}
		userRate, err := limiter.NewRateFromFormatted(telegramConf.UserRatelimit)
		if err != nil {
			panic(err)
		}
		tl.UserLimiter = limiter.New(store, userRate)
	}

	tl.ExcludeKeyMap = make(map[string]struct{}, len(telegramConf.ExcludeKeys))
	for _, k := range telegramConf.ExcludeKeys {
		tl.ExcludeKeyMap[k] = struct{}{}
	}

	return tl
}

const (
	// telegram bot 系统请求限制key
	TelegramGeneralLimiterKey = "telegram_bot_general_request_limit"
	// telegram bot 用户请求限制key前缀，防止userID和系统key冲突
	TelegramUserLimiterPrefix = "telegram_user_limiter_prefix"
)

func (tLimiter *TelegramLimiter) ReachedLimit(userID string) bool {
	// 判断 userID 是否在 exclude key中
	if _, ok := tLimiter.ExcludeKeyMap[userID]; ok {
		return false
	}

	// 1. 判断用户请求是否超限
	userRes := tLimiter.userReachedLimit(userID)
	// 2. 判断系统请求是否超限
	// 步骤1 和 2都要运行一次，只有运行了，系统底层才会统计使用次数
	teleRes := tLimiter.teleReachedLimit()

	// 综合返回结果
	return userRes || teleRes
}

// userReachedLimit 判断 userID 是否达到上线
// true 说明达到上限，需要进行限制
// false 说明没有达到，直接放行
func (tLimiter *TelegramLimiter) userReachedLimit(userID string) bool {
	// 如果限流器为nil，说明没有配置限流，默认不限流
	if tLimiter.UserLimiter == nil {
		return false
	}
	if _, ok := tLimiter.ExcludeKeyMap[userID]; ok {
		return false
	}

	newKey := fmt.Sprintf("%s_%s", TelegramUserLimiterPrefix, userID)
	userCtx, err := tLimiter.UserLimiter.Get(context.TODO(), newKey)
	if err != nil {
		log.Errorf("UserLimiter Get key=%s error: %v", newKey, err)
		return true
	}
	return userCtx.Reached
}

// teleReachedLimit 判断telegram总请求数是否达到上线
// true 说明达到上限，需要进行限制
// false 说明没有达到，直接放行
func (tLimiter *TelegramLimiter) teleReachedLimit() bool {
	// 如果限流器为nil，说明没有配置限流，默认不限流
	if tLimiter.Limiter == nil {
		return false
	}

	generalCtx, err := tLimiter.Limiter.Get(context.TODO(), TelegramGeneralLimiterKey)
	if err != nil {
		log.Errorf("Limiter Get key=%s error: %v", TelegramGeneralLimiterKey, err)
		return true
	}
	return generalCtx.Reached
}

// NewTelegramBot 根据 token 生成默认的 telegram bot
func NewTelegramBot(token string, proxyUrl string) *tgbotapi.BotAPI {
	// 判断是否使用代理http client
	var HTTPClient *http.Client
	if proxyUrl != "" {
		urlObj, err := url.Parse(proxyUrl)
		if err != nil {
			log.Errorf("parse proxy_url %s error: %v", proxyUrl, err)
			panic(err)
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
		log.Error(err)
		panic(err)
	}
	log.Infof("Authorized on account %s\n", bot.Self.UserName)

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
		// lib.PrintJson("DEBUG-syncProcessTelegramCommand", update)
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		msg.ReplyToMessageID = update.Message.MessageID

		fromUser := update.Message.From.UserName
		if !update.Message.IsCommand() { // ignore any non-command Messages
			s.log.Warnf("message from user %s is not command, msg=%s, skip",
				fromUser, update.Message.Text)
			msg.Text = `please use "/chat youar_message" to chat with me by openai`
			s.telegramReplayMessage(&msg)
			continue
		}

		// 初始化返回消息结构体
		commandDetail := lib.CompressMessage(update.Message.CommandArguments())
		if len(commandDetail) == 0 {
			msg.Text = `"/chat youar_message" youar_message is empty, skip`
			s.telegramReplayMessage(&msg)
			continue
		}
		req := new(pb.SummaryReuqest)
		req.PromptDetail = commandDetail
		start := time.Now()

		// 开始判断是否超过请求限制
		if s.TelegramLimiter.ReachedLimit(fromUser) {
			msg.Text = "您今天超过使用次数限制，请明天再来"
			s.telegramReplayMessage(&msg)
			continue
		}

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
		msg.Text = fmt.Sprintf("%s\n\n耗时: %s", msg.Text, costTime.String())
		s.telegramReplayMessage(&msg)
	}
}

// telegramReplayDefaultMessage telegram 机器人返回默认消息
func (s *AdminService) telegramReplayMessage(message *tgbotapi.MessageConfig) {
	if message == nil {
		return
	}
	if _, err := s.TelegramBot.Send(message); err != nil {
		s.log.Errorf("[telegramReplayMessage] send telegram msg %v, error: %v", message, err)
	}
}

var ErrNotSupportTelegramCommand = errors.New("not support telegram bot command")
