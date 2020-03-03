package main

import (
	"fmt"
	"os"
)

func main() {
	a := os.Args
	b := os.Args[1:]
	c := os.Args[3]
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

// 命令行传参数
// go build main.go
// ./main h j j d
//以下为执行结果
// [./main h j j d]
// [h j j d]
// j
