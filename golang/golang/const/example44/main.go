package main

import (
	"fmt"
	"runtime"
)

var (
	flag = false
	str  string
)

func foo() {
	flag = true
	str = "setup complete!"
}

func main() {
	runtime.GOMAXPROCS(1) //设置单核
	go foo()
	for {
		if flag {
			break
		}
	}
	fmt.Println(str)
}
