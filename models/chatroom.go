package models

import (
	"time"
//_ "database/sql"
//_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"fmt"
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
	onlineMembers []uint32
}

type ChatServer struct {
	rooms      map[uint32]*ChatRoom
	idReqChan  chan string
	idRespChan chan uint32
	dispatcher chan []byte
}

func init() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
}

func NewChatServer() *ChatServer {
	chatServer := &ChatServer{
		rooms: make(map[uint32]*ChatRoom),
		idReqChan: make(chan string, 1),
		idRespChan: make(chan uint32, 1),
		dispatcher: make(chan []byte, 1),
	}
	go runIdServer(chatServer)
	go runDispatcher(chatServer)
	return chatServer
}

func (this *ChatServer) NewChatRoom(title string, chatType byte, userIdList []uint32) {
	id := this.newRoomId()
	online := make([]uint32)
	for userId := range userIdList {
		s := App.sessionServer.GetSessionByUserId(userId)
		if s.conn {
			online = append(online, userId)
		}
	}
	c := &ChatRoom{id, title, chatType, time.Now(), online}
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

func runDispatcher(chatServer *ChatServer) {
	var msg Message
	for {
		data := <-chatServer.dispatcher
		if err := json.Unmarshal(data, &msg); err != nil {
			fmt.Printf("json decode failed: %s\n", data)
			continue
		}
		switch msg.targetType {
		case msgTargetUser:
			sendMessage(msg.targetId, msg, data)
		case msgTargetGroup:
			m, _ := chatServer.rooms[msg.targetId]
			if m {
				for userId := range m.onlineMembers {
					sendMessage(userId, msg, data)
				}
			}
		default:
			fmt.Printf("error message type: %s\n", msg.targetType)
		}
	}
}

func sendMessage(targetId uint32, msg Message, data []byte) {
	s := App.sessionServer.GetSessionByUserId(targetId)
	if s {
		select {
		case s.conn.send <- data:
		default:
			close(s.conn.send)
			App.sessionServer.DelSession(s.sid)
		}
	}
}