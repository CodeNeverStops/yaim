package models

import (
//"fmt"
//_ "database/sql"
//_ "github.com/go-sql-driver/mysql"
)

const (
	msgTypeText = 0
	msgTypeImage = 1
	msgTypeVoice = 2

	msgTargetUser = 0
	msgTargetGroup = 1
)

type Message struct {
	senderId    uint32  `json:"sender_id"`
	targetId    uint32    `json:"target_id"`
	targetType  byte    `json:"target_type"`
	messageType byte    `json:"message_type"`
	message     []byte    `json:"message"`
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
