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
		defer fmt.Println("A is defer")
		func() {
			defer fmt.Println("B is defer")
			runtime.Goexit() //退出当前goroutine，defer后的还是会执行，而且是先进后出
			fmt.Println("B") //不会执行
		}()
		fmt.Println("A") //不会执行
	}()
	wg.Wait()
}
