package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name  string
	age   int
	score float32
}

func test(b interface{}) {
	a := reflect.TypeOf(b)
	fmt.Println(a)
	v := reflect.ValueOf(b)
	fmt.Println(v)
	k := v.Kind()
	fmt.Println(k)
	f := v.Interface()
	fmt.Println(f)
	stu, ok := f.(student)
	if ok {
		fmt.Printf("%v%T\n", stu, stu)
	}
}
func main() {
	var k student = student{ //反射获取结构体
		name:  "stu01",
		age:   18,
		score: 92,
	}
	test(k)
}
