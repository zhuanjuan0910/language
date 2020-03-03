package main

import "fmt"

const (
	a = iota //特殊常量，从零开始，可自增
	b        //1
	c        //2
	d        //3
	e        //4
	f        //5
)

func main() {
	fmt.Println(a, b, c, d, e, f)
}
