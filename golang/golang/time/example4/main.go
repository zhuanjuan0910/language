package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTicker(3 * time.Second) //三秒执行一次
	for v := range t1.C {                 //C <-chan Time // The channel on which the ticks are delivered.
		fmt.Println("hello, ", v)
	}
}

//定时器的作用
