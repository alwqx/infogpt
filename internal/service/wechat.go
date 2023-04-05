package service

import (
	"infogpt/internal/conf"

	wechat "github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
)

func NewOfficialAccount(adminConf *conf.Admin) *officialaccount.OfficialAccount {
	wc := wechat.NewWechat()
	// 这里本地内存保存 access_token
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:          adminConf.OfficialAccount.AppId,
		AppSecret:      adminConf.OfficialAccount.AppSecret,
		Token:          adminConf.OfficialAccount.Token,
		EncodingAESKey: adminConf.OfficialAccount.EncodingAesKey, // 控制是否加密
		Cache:          memory,
	}
	return wc.GetOfficialAccount(cfg)
}
