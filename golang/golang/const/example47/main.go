package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}
func read(ch chan int) {
	for {
		var b int
		b = <-ch
		fmt.Println(b)
		time.Sleep(time.Second)
	}
}
func main() {
	inchan := make(chan int, 5)
	go write(inchan) //多个goroutine同时访问不需要加锁，其内部实现了同步
	go read(inchan)
	time.Sleep(10 * time.Second)
}
