package main

import (
	"fmt"
	"reflect"
)

func main() {
	var b int = 1
	b = 200
	testInt(&b)
	fmt.Println(b)
}
func testInt(b interface{}) { //用反射来改变指针指向的值
	v := reflect.ValueOf(b)
	fmt.Println(v)
	v.Elem().SetInt(100)
	c := v.Elem().Int()
	fmt.Printf("get value  interface{} %d\n", c)
	fmt.Printf("string val:%d\n", v.Elem().Int())
}
