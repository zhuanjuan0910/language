package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var lock sync.Mutex

func testLock() {
	var a map[int]int
	a = make(map[int]int, 5)
	var count int32
	a[8] = 10
	a[3] = 10
	a[2] = 10
	a[1] = 10
	a[18] = 10
	for i := 0; i < 2; i++ {
		go func(b map[int]int) {

			lock.Lock()
			b[8] = rand.Intn(100)
			time.Sleep(10 * time.Microsecond) //微妙

			lock.Unlock()

		}(a)
	}
	for i := 0; i < 100; i++ {
		go func(b map[int]int) {
			for {

				lock.Lock()
				time.Sleep(time.Millisecond)

				lock.Unlock()
				atomic.AddInt32(&count, 1)
			}
		}(a)
	}
	time.Sleep(time.Second * 20)
	fmt.Println(atomic.LoadInt32(&count))
}
func main() {
	testLock()

}
