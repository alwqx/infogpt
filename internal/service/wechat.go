package service

import (
	pb "infogpt/api/admin/v1"
	"infogpt/internal/conf"
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
		} else {
			chatReq := &pb.OpenaiChatReuqest{
				Message: msg.Content,
			}
			// 如果消息过于复杂，OpenAI处理时间超过5秒，微信会断开连接，并且重试3次
			// 这里要做好检查，如果超过5秒，就返回success或者提示信息
			// 然后把问题缓存起来，新起一个逻辑，去请求答案
			// 详情见: https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Passive_user_reply_message.html
			chatResp, err := s.OpenaiChat(ctx, chatReq)
			if err != nil {
				s.log.Errorf("[processOfficialAccountMessage] serveOfficialWechat error: %v", err)
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
		s.log.Errorf("[processOfficialAccountMessage] serve.Serve error: %v", err)
		return
	}
	err = server.Send()
	if err != nil {
		s.log.Errorf("[processOfficialAccountMessage] server.Send error: %v", err)
	}
}
