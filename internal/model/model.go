package model

import "time"

// WeChatMessageCacheItem 用于缓存的微信公众号聊天消息
type WeChatMessageCacheItem struct {
	Message    string    // 原始消息压缩后
	Replay     string    // openai返回的消息
	Status     string    // 消息的处理进度
	ChatTime   time.Time // 聊天发生的时间
	ReplayTime time.Time // openai返回时间
}

const (
	WechatMessageStatusInit    = "init"    // 微信消息状态：刚初始化
	WechatMessageStatusRequest = "request" // 微信消息状态：刚初始化
	WechatMessageStatusError   = "error"   // 微信消息状态：openai返回，完成
	WechatMessageStatusDone    = "done"    // 微信消息状态：openai返回，完成
)
