package main

import "fmt"

func main() {
	test()
}
func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()
	panic("hello world")
}
