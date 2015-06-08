package models

import (
	"time"
	//"fmt"
	//_ "database/sql"
	//_ "github.com/go-sql-driver/mysql"
)

type Message struct {
	chatId      string // ug|123|456, uu|123|789
	senderId    int
	message     string
	messageType string
	createAt    time.Time
}

func init() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
}

func saveOne(msg *Message) bool {
	return true
}

func getHistoryAll(chatId string) []Message {
	return []Message{}
}
