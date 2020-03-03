package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var sum int32 = 100
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.CompareAndSwapInt32(&sum, 100, sum+1) //sum的值等于100就赋值新值sum+1,交换并赋值
		}()
	}
	wg.Wait()
	fmt.Println(sum) //101
}
