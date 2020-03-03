package main

import "fmt"

const ( //iota自增被打断必须显示恢复
	a = iota
	b
	c = "hello"
	d
	e
	f = iota //显示恢复，自增数c，d，e也要计数
	g
)

func main() {
	fmt.Println(a, b, c, d, e, f, g) //0,1,hello,hello,hello,5,6
}
