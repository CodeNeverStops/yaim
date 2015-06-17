package models

import (
	"time"
	"strconv"
	"math/rand"
	"crypto/sha1"
	"fmt"
)

type session struct {
	sid     string
	userId  uint32
	account string
	conn    *connection
}

type SessionPool struct {
	sid map[string]*session
	uid map[uint32]*session
}

var UserSession map[uint32]*session

func NewSessionPool() *SessionPool {
	return &SessionPool{
		sid: make(map[string]*session),
		uid: make(map[uint32]*session),
	}
}

func (this *SessionPool) NewSession(userId uint32, account string, conn *connection) string {
	sid := this.genSessionId(userId, account)
	s := &session{sid, userId, account, conn}
	this.addSession(s)
	return sid
}

func (this *SessionPool) GetSessionByUserId(userId uint32) *session {
	s, _ := this.uid[userId]
	return s
}

func (this *SessionPool) GetSession(sid string) *session {
	s, _ := this.sid[sid]
	return s
}

func (this *SessionPool) DelSession(sid string) {
	s := this.GetSession(sid)
	if s {
		delete(this.sid, sid)
		delete(this.uid, s.userId)
	}
}

func (this *SessionPool) addSession(s *session) {
	this.sid[s.sid] = s;
	this.uid[s.userId] = s
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