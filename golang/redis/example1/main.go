package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	con, err := redis.Dial("tcp", "ip：端口") //建立连接
	if err != nil {
		fmt.Println("dial failed,err")
		return
	}
	fmt.Println("dial success,")
	defer con.Close() //关闭连接
	_, err = con.Do("auth", "密码")
	_, err = con.Do("set", "abc", 100)
	if err != nil {
		fmt.Println("set failed", err)
		return
	}
	r, err := redis.Int(con.Do("get", "abc"))
	if err != nil {
		fmt.Println("get failed", err)
		return
	}
	fmt.Println(r)

}
