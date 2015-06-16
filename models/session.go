package models

import (
	"time"
	"strconv"
	"math/rand"
	"crypto/sha1"
	"fmt"
)

type session struct {
	sessionId string
	userId    uint32
	account   string
	conn      *connection
}

type SessionPool struct {
	pool map[uint32]*session
}

func NewSessionPool() *SessionPool {
	return &SessionPool{
		pool: make(map[uint32]*session),
	}
}

func (this *SessionPool) NewSession(userId uint32, account string, conn *connection) {
	sessionId := this.genSessionId(userId, account)
	s = &session{sessionId, userId, account, conn}
	this.addSession(s)
}

func (this *SessionPool) GetSession(userId uint32) *session {
	s, _ := this.pool[userId]
	return s
}

func (this *SessionPool) DelSession(userId uint32) {
	delete(this.pool, userId)
}

func (this *SessionPool) addSession(s *session) {
	this.pool[s.userId] = s;
}

func (this *SessionPool) genSessionId(userId uint32, account string) string {
	secureKey := "^&`d1U(_]0?"
	now := time.Now().Unix()
	data := []byte{}
	data = strconv.AppendInt(data, rand.Int63n(now), 10)
	data = strconv.AppendUint(data, userId, 10)
	data = strconv.AppendQuote(data, secureKey)
	data = strconv.AppendInt(data, now, 10)
	return fmt.Sprintf("%X", sha1.Sum(data))
}