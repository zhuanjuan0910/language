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
	database, err := sqlx.Open("mysql", "用户名：密码@tcp(host:端口号)/数据库名") //更新数据库
	if err != nil {
		fmt.Println("open database is failed", err)
		return
	}
	Db = database
}
func main() {
	r, err := Db.Exec("update person set username=? where user_id=?", "stu003", 2)
	if err != nil {
		fmt.Println("update failed", err)
		return
	}
	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("row failed", err)
		return
	}
	fmt.Println("update success", row)

}
