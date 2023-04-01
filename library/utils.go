package library

import (
	"encoding/json"
	"log"
	"net/url"
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
