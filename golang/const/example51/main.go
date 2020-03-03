package main

import (
	"fmt"
	"math/rand"
	"time"
)

func newTest() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano()) //设置随机数种子，加上这行代码，可以保证每次随机都是随机的
	go func() {
		c <- rand.Int()
	}()
	return c
}

func main() {
	t := newTest()
	fmt.Println(<-t) //等待goroutine结束返回
}
