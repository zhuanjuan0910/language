package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

var a interface{} = nil         // tab = nil, data = nil 只有tab和data都等于nil时，接口才等于nil
var b interface{} = (*int)(nil) // tab 包含 *int 类型信息, data = nil

type iface struct {
	itab, data uintptr
}

func main() {
	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))

	fmt.Println(a == nil, ia)
	fmt.Println(b == nil, ib, reflect.ValueOf(b).IsNil())
}
