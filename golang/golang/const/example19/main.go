package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var ch chan int = make(chan int, 2)
	close(ch) //不能忘已经关闭的通道写东西
	ch <- 2

}
