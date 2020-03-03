package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var rwLock sync.RWMutex

func testRwlock() {
	a := make(map[int]int, 5)
	a[1] = 10
	a[2] = 10
	a[3] = 10
	a[4] = 10
	a[5] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {
			rwLock.Lock()
			b[4] = rand.Intn(100)
			rwLock.Unlock()
		}(a)
	}
	for i := 0; i < 10; i++ {
		go func(b map[int]int) {
			rwLock.RLock() //读锁
			fmt.Println(a)
			rwLock.RUnlock()
		}(a)
		time.Sleep(2 * time.Second)
	}

}
func main() {
	testRwlock() //读多写少时用读锁
}
