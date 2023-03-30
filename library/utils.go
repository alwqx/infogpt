package library

import (
	"encoding/json"
	"log"
)

func PrintJson(title string, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Printf("Marshal %v error: %v\n", v, err)
		return
	}
	log.Printf("[PrintJson][%s] %s\n", title, string(b))
}
