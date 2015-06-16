package models

import (
	"time"
//"fmt"
//_ "database/sql"
//_ "github.com/go-sql-driver/mysql"
	"crypto"
)

const (
	chatTypeUser byte = 'u'
	chatTypeGroup byte = 'g'
)

const (
	cmdIncr string = "incr"
	cmdShutdown string = "shutdown"
)

type ChatRoom struct {
	id            uint32
	title         string
	chatType      byte // u or g
	createAt      time.Time
	onlineMembers map[uint32]*User
	memberIdList  []uint32
}

type ChatServer struct {
	rooms      map[uint32]*ChatRoom
	idReqChan  chan string
	idRespChan chan string
}

func init() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
}

func NewChatServer() {
	chatServer := &ChatServer{
		rooms: make(map[uint32]*ChatRoom),
		idReqChan: make(chan string, 1),
		idRespChan: make(chan uint32, 1),
	}
	go runIdServer(chatServer)
	return chatServer
}

func (this *ChatServer) NewChatRoom(title string, chatType byte, userIdList []uint32) {
	id := this.newRoomId()
	// @todo filter online members out
	c := &ChatRoom{id, title, chatType, time.Now(), make(map[uint32]*User), userIdList}
	this.addChatRoom(c)
}

func (this *ChatServer) addChatRoom(room *ChatRoom) {
	this.rooms[room.id] = room;
}

func (this *ChatServer) newRoomId() uint32 {
	this.idReqChan <- cmdIncr
	return <-this.idRespChan
}

func runIdServer(chatServer *ChatServer) {
	var id uint32 = 0
	for {
		msg := <-chatServer.idReqChan
		switch msg {
		case cmdIncr:
			id++
			chatServer.idRespChan <- id
		case cmdShutdown:
		// @todo save chat id
		}
	}
}