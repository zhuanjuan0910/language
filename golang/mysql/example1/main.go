package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //引入两个外部包
	"github.com/jmoiron/sqlx"
)

type person struct {
	UserId   int    `db:"user_id"`
	UserName string `db:"username"` //flag里别忘记双引号
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
	r
	database, err := sqlx.Open("mysql", "用户名：密码@tcp(host:端口号)/数据库名") //和数据库建立连接
	if err != nil {
		fmt.Println("open database is failed", err)
		return
	}
	Db = database
}
func main() {
	r, err := Db.Exec("insert into person(username,sex,email)values(?,?,?)", "stu001", "man", "stu01@qq.com") //执行插入语句
	if err != nil {
		fmt.Println("insert failed", err)
		return
	}
	id, err := r.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(id)

}
