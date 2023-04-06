package library

import (
	"encoding/json"
	"log"
	"net/url"
	"strings"
)

func PrintJson(title string, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("Marshal %v error: %v\n", v, err)
		return
	}
	log.Printf("[PrintJson][%s] %s\n", title, string(b))
}

// CheckUrl 检查url链接是否正常
func CheckUrl(urlStr string) error {
	_, err := url.Parse(urlStr)
	return err
}

// CompressMessage 压缩消息
// 当前主要是去掉首尾空格
func CompressMessage(msg string) string {
	return strings.TrimSpace(msg)
}
