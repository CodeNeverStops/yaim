package models
import "time"

type Session struct {
	sessionId string
	account string
	activeTime time.Time
}
