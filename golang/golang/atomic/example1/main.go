package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var sum int32

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//sum += 1
			atomic.AddInt32(&sum, 1) //原子性操作保证只有一个goroution在操作
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}

// 原子增、减值
// 用于对变量值进行原子增操作，并返回增加后的值。

// 第一个参数值必须是一个指针类型的值，以便施加特殊的CPU指令。
// 第二个参数值的类型和第一个被操作值的类型总是相同的。
