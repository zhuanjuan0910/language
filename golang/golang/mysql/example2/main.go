package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type person struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}
type place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "用户名：密码@tcp(host:端口号)/数据库名")
	if err != nil {
		fmt.Println("open database is failed", err)
		return
	}cd
	Db = database

}
func main() {
	var person []person
	err := Db.Select(&person, "select * from person where user_id=?", 2) //查询数据库
	if err != nil {
		fmt.Println("select failed", err)
		return
	}
	fmt.Println(person)
}
