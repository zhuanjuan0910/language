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
	r, err := Db.Exec("delete from person where user_id=?", 3) //删除表中的字段
	if err != nil {
		fmt.Println("delete failed", err)
		return

	}
	row, err := r.RowsAffected()
	if err != nil {
		fmt.Println("row failed", err)
	}
	fmt.Println("delete success", row)
}
