package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	con, err := redis.Dial("tcp", "ip：端口")
	if err != nil {
		fmt.Println("con is failed", err)
		return
	}
	defer con.Close()
	_, err = con.Do("auth", "密码")
	if err != nil {
		fmt.Println("yanzhengshibai", err)
		return
	}
	_, err = con.Do("hmset", "student", "name", "zhangsan", "age", 16, "sex", "man") //hash存储，特别适合存对象
	if err != nil {
		fmt.Println("hset is failed", err)
		return

	}
	r, err := redis.Strings(con.Do("hmget", "student", "name", "age", "sex"))
	if err != nil {
		fmt.Println("hget is failed", err)
		return
	}
	for _, v := range r {
		fmt.Println(v)
	}
}
