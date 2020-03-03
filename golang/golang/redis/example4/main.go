package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	con, err := redis.Dial("tcp", "ip：端口")
	if err != nil {
		fmt.Println("con failed", err)
		return
	}
	defer con.Close()
	_, err = con.Do("auth", "密码")
	if err != nil {
		fmt.Println("yanzhengshibai", err)
	}
	_, err = con.Do("lpush", "book_list", "abc", "def", 290) //向列表的头部插入数据
	if err != nil {

		fmt.Println("lpush is failed", err)
	}
	r, err := redis.String(con.Do("lpop", "book_list")) //移出并返回列表的第一个数据
	if err != nil {
		fmt.Println("lpop is failed", err)
		return
	}
	fmt.Println(r)
}
