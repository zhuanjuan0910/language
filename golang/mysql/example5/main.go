package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "用户名：密码@tcp(host:端口号)/数据库名")
	if err != nil {
		fmt.Println("open database failed", err)
		return
	}
	Db = database
}
func main() {
	con, err := Db.Begin() //开启事务
	if err != nil {
		fmt.Println("begin failed", err)
		return
	}
	r, err := con.Exec("insert into person(username,sex,email)values(?,?,?)", "stu005", "man", "stu005@qq.com") //注意这里是事务的exec方法
	if err != nil {
		fmt.Println("exec failed", err)
		con.Rollback() //事务回滚
		return
	}
	row, err := r.LastInsertId()
	if err != nil {
		fmt.Println("row failed", err)
		con.Rollback()
		return
	}
	fmt.Println("row sucess", row)
	con.Commit() //事务提交
}
