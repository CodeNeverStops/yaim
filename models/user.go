package yaim

import (
	"fmt"
	//_ "database/sql"
	//_ "github.com/go-sql-driver/mysql"
)

type User struct {
	account string
}

func init() {
	//db, err := sql.Open("mysql", "user:password@/dbname")
}

func signup(account string, password string) bool {
	fmt.Printf("account:%s,password:%s", account, password)
	fmt.Println("=========")
	return true
}

func signin(account string, password string) bool {
	fmt.Printf("account:%s,password:%s", account, password)
	fmt.Println("=========")
	return true
}
