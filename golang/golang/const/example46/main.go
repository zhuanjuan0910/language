package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			fmt.Println(i)
			runtime.Gosched() //将当前 goroutine 暂停，放回队列等待下次被调度执行。所以每次的执行结果都不一样
		}
	}()
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("hello world")
		}()

	}
	wg.Wait()

}
