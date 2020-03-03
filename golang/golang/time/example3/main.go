package main

import (
	"fmt"
	"time"
)

//测试app运行时间
func test() {
	t1 := time.Now()
	for i := 0; i < 10000; i++ {
		time.Sleep(2 * time.Nanosecond)
		continue
	}
	running := time.Since(t1)
	fmt.Println("app is running", running)
}
func main() {
	test()
}
