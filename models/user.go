package models

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	account string
}

var dsn string
//var db sql.DB

func init() {
	fmt.Println("user init")
	dsn = "user:1u2s3e4r5@tcp(10.251.18.39:3306)/test2?charset=utf8"
}

func SignUp(account string, password string, passwordRetry string) bool {
	fmt.Printf("account:%s,password:%s,passwordRetry:%s", account, password, passwordRetry)
	fmt.Println("=========")

	if password != passwordRetry {
		panic("password is not same")
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("connect to database failed:" + err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("insert into users set account=?, password=?")
	if err != nil {
		panic("prepare failed:" + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(account, password)
	if err != nil {
		panic(err.Error())
	}

	return true
}

func SignIn(account string, password string) bool {
	fmt.Printf("account:%s,password:%s", account, password)
	fmt.Println("=========")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic("connect to database failed:" + err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare("select password from users where account=?")
	if err != nil {
		panic("prepare failed:" + err.Error())
	}
	defer stmt.Close()

	var userPassword string
	err = stmt.QueryRow(account).Scan(&userPassword)
	if err != nil {
		panic(err.Error())
	}

	if password == userPassword {
		return true
	}

	return false
}
