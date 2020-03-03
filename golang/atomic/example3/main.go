package main

import (
	"fmt"
	"sync/atomic"
)

//第一个参数都为指针类型
func main() {
	var a int32 = 20
	atomic.StoreInt32(&a, 83) //设置值
	fmt.Println(a)
	v := atomic.LoadInt32(&a) //读取值
	fmt.Println(v)
	var n int32 = 300
	var o int32 = 27
	atomic.SwapInt32(&o, n) //把新值交换给旧值，并返回旧值
	fmt.Println(o, n)
}
