package service

import (
	"context"
	"fmt"
	"time"

	pb "infogpt/api/admin/v1"
	"infogpt/internal/conf"
	"infogpt/internal/model"
	lib "infogpt/library"

	"github.com/gin-gonic/gin"
	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/message"
)

func NewOfficialAccount(adminConf *conf.Admin) *officialaccount.OfficialAccount {
	wc := wechat.NewWechat()
	// 这里本地内存保存 access_token
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:          adminConf.Wechat.AppId,
		AppSecret:      adminConf.Wechat.AppSecret,
		Token:          adminConf.Wechat.Token,
		EncodingAESKey: adminConf.Wechat.EncodingAesKey, // 控制是否加密
		Cache:          memory,
	}
	return wc.GetOfficialAccount(cfg)
}

// isOfficialAccountEnable 根据配置信息判断是否开启公众号功能
func (s *AdminService) isOfficialAccountEnable(officialAccountConf *conf.WeChat) bool {
	if officialAccountConf == nil {
		return false
	}
	if officialAccountConf.AppId == "" ||
		officialAccountConf.AppSecret == "" ||
		officialAccountConf.Token == "" {
		return false
	}

	// EncodingAesKey 可以不提供，如果提供长度必须是43
	keyLen := len(officialAccountConf.EncodingAesKey)
	if keyLen != 0 && keyLen != lib.WeChatOfficialAccountEncodingAesKeyLen {
		s.log.Warnf("[isOfficialAccountEnable] encoding_aes_key len shoud be 43 but get %d", keyLen)
		return false
	}

	return true
}

// ProcessOfficialAccountMessage 处理微信公众号消息
func (s *AdminService) ProcessOfficialAccountMessage(ctx *gin.Context) {
	if !s.EnableWeChat {
		s.log.Warn("[ProcessOfficialAccountMessage] not enable wechat official account, skip")
		return
	}
	server := s.OfficialAccount.GetServer(ctx.Request, ctx.Writer)
	server.SetMessageHandler(func(msg *message.MixMessage) *message.Reply {
		// 先判断是否是关键字自动回复
		respText := new(message.Text)
		if replyInfo, ok := s.AdminConf.Wechat.AutoReplay[msg.Content]; ok {
			respText.Content = message.CDATA(replyInfo)
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: respText}
		}

		// 判断缓存中是否存在消息
		msgKey := lib.CompressMessage(msg.Content)
		reply, ok := s.getWechatMessageFromCache(msgKey)
		// if ok && reply.Status == model.WechatMessageStatusDone {
		if ok {
			respText.Content = message.CDATA(s.truncateWechatMessage(reply.Replay))
			return &message.Reply{MsgType: message.MsgTypeText, MsgData: respText}
		}

		// 不存在，则开启goroutine请求reply
		respCh := make(chan string)
		go s.getAndCacheWechatMessageFromOpenAI(msgKey)
		go func() {
			// 这个地方不能无限重试，需要做判断
			sleepTime := 900 * time.Millisecond
			cnt := 1
			timeout := 5 * time.Minute
			for {
				item, ok := s.getWechatMessageFromCache(msgKey)
				if !ok || item.Status != model.WechatMessageStatusDone {
					s.log.Warnf("[ProcessOfficialAccountMessage] not get replay of %s, sleep", msgKey)
					time.Sleep(sleepTime)
					cnt++
					if time.Duration(cnt)*sleepTime > timeout {
						respCh <- "请求超时"
						break
					}
					continue
				}
				reply := fmt.Sprintf("%s\n\n耗时: %s",
					item.Replay, item.ReplayTime.Sub(item.ChatTime).String())
				respCh <- reply
				s.log.Warnf("[ProcessOfficialAccountMessage] get replay of %s, sleep", msgKey)
				break
			}
		}()
		select {
		case reply := <-respCh:
			respText.Content = message.CDATA(s.truncateWechatMessage(reply))
		case <-time.After(4500 * time.Millisecond):
			// 超时，返回默认信息提示用户
			respText.Content = message.CDATA("您的消息耗时较长,已经缓存后台处理,请等待30秒左右输入相同的问题")
		}

		// 如果返回消息过长，超过微信限制，会被拒绝发送，这里加上判断，如果过长，则分割发送
		return &message.Reply{MsgType: message.MsgTypeText, MsgData: respText}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		s.log.Errorf("[processOfficialAccountMessage] serve.Serve error: %v", err)
		return
	}
	err = server.Send()
	if err != nil {
		s.log.Errorf("[processOfficialAccountMessage] server.Send error: %v", err)
	}
}

// getWechatMessageFromCache
func (s *AdminService) getWechatMessageFromCache(key string) (*model.WeChatMessageCacheItem, bool) {
	item, ok := s.WeChatMessageCache.Get(key)
	if !ok {
		return nil, false
	}

	res, ok := item.(*model.WeChatMessageCacheItem)
	if !ok {
		return nil, false
	}
	return res, true
}

// getAndCacheWechatMessageFromOpenAI
func (s *AdminService) getAndCacheWechatMessageFromOpenAI(msg string) {
	item := new(model.WeChatMessageCacheItem)
	item.Message = msg
	item.ChatTime = time.Now()
	s.log.Warnf("[getAndCacheWechatMessageFromOpenAI] finish replay %s, start=%s\n", msg, item.ChatTime.String())
	chatReq := &pb.OpenaiChatReuqest{
		Message: msg,
	}
	item.Status = model.WechatMessageStatusRequest
	s.WeChatMessageCache.SetDefault(item.Message, item)

	chatResp, err := s.OpenaiChat(context.Background(), chatReq)
	if err != nil {
		s.log.Errorf("[getAndCacheWechatMessageFromOpenAI] OpenaiChat error: %v", err)
		item.Replay = err.Error()
		item.Status = model.WechatMessageStatusError
	} else {
		item.Replay = chatResp.Message
		item.Status = model.WechatMessageStatusDone
	}
	item.ReplayTime = time.Now()
	s.WeChatMessageCache.SetDefault(item.Message, item)
	s.log.Warnf("[getAndCacheWechatMessageFromOpenAI] finish replay %s, start=%s, end=%s\n",
		msg, item.ChatTime.String(), item.ReplayTime.String())
}

// truncateWechatMessage 微信返回消息有长度限制，超过则截断
func (s *AdminService) truncateWechatMessage(input string) string {
	rs := []rune(input)
	if len(rs) < lib.WeChatReplayMessageLen {
		return input
	}

	res := rs[:lib.WeChatReplayMessageLen]
	return string(res)
}
