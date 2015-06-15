package models

import (
	"time"
	//"fmt"
	//_ "database/sql"
	//_ "github.com/go-sql-driver/mysql"
	"debug/dwarf"
)

type Chatroom struct {
	id       int
	chatType byte // u or g
	createAt time.Time
	members  map[int]*User
	sessions map[int]*session
}

func init() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
}
