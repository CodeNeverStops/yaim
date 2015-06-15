package models
import "time"

type session struct {
	sessionId string
	account string
	conn *connection
}

var sessionPool map[int]session

func (s *session) new(account string, conn *connection) {



}

func genSessionId() {

}
