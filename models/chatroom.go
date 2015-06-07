package yaim

import (
	"time"
	//"fmt"
	//_ "database/sql"
	//_ "github.com/go-sql-driver/mysql"
)

type Chatroom struct {
	id       int
	chatType string
	createAt time.Time
	members  []User
}

func init() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
}
