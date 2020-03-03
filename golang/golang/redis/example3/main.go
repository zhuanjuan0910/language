package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	con, err := redis.Dial("tcp", "ip：端口")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	defer con.Close()
	_, err = con.Do("auth", "密码")
	_, err = con.Do("expire", "abc", 10) //设置过期时间
	if err != nil {
		fmt.Println("设置失败", err)
		return
	}
	fmt.Println("设置成功")
}
