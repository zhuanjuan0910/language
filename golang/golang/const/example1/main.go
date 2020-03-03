package main

import "fmt"

const (
	a = "abc"
	c //   c="abc" 当常量未给定类型和值时，值与上一个常量相同
)

func main() {
	fmt.Println(a, c)
}
