package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	con, err := redis.Dial("tcp", "ip：端口")
	if err != nil {
		fmt.Println("dial failed", err)
		return
	}
	defer con.Close()
	_, err = con.Do("auth", "密码")
	if err != nil {
		fmt.Println("密码验证失败")
	}
	_, err = con.Do("mset", "abc", 100, "def", 200) //string批量操作
	if err != nil {
		fmt.Println("mset failed", err)
		return
	}
	data, err := redis.Ints(con.Do("mget", "abc", "def")) //批量获取string的value
	if err != nil {
		fmt.Println("mget failed", err)
		return
	}
	for _, v := range data {
		fmt.Println(v)
	}
}
