package main

import "fmt"

func main() {
	data := make(chan int, 3) //缓存区可以存3个元素
	exit := make(chan bool)
	data <- 1
	data <- 2
	data <- 3
	go func() {
		for d := range data {
			fmt.Println(d)
		}
		exit <- false
	}()
	data <- 4 //如果缓存区已满，阻塞
	data <- 5
	close(data)
	<-exit

}
