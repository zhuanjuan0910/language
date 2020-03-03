package main

import (
	"fmt"
	"unsafe"
)

const ( //常量可以是len，cap，unsafe.Sizeof等编译期可确定结果的函数返回值
	a = "hello world"
	b = len(a)
	c = unsafe.Sizeof(b) //返回字节数
)

func main() {
	fmt.Println(a, b, c)
}
