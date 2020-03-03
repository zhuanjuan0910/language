package main

import (
	"fmt"
)

func add(a, b int) (c int) {
	c = a + b
	return //return语句没有参数时返回当前各变量的值
}

func calc(a, b int) (sum int, avg int) {
	sum = a + b
	avg = (a + b) / 2

	return
}

func main() {
	var a, b int = 1, 2
	c := add(a, b)
	sum, avg := calc(a, b)
	fmt.Println(a, b, c, sum, avg)
}
