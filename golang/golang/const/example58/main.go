package main

import (
	"fmt"
	"reflect"
)

type student struct {
	age   int
	name  string
	score float32
}

func (s student) set(age int, name string, score float32) {
	s.age = age
	s.name = name
	s.score = score
}
func (s student) Print() {
	fmt.Println(s)
}
func testStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	fmt.Println(val, kd)
	if kd != reflect.Struct {
		fmt.Println("kd is not struct")
		return
	}
	//获取字段数量
	fleds := val.NumField()
	fmt.Printf("val has %d fleds", fleds)
	//获取字段的类型
	for i := 0; i < fleds; i++ {
		fmt.Printf("%d%v\n", i, val.Field(i).Kind())
	}
	//获取方法数量
	methods := val.NumMethod()
	fmt.Printf("val has %d method", methods)
	//反射调用print方法
	var parms []reflect.Value
	val.Method(0).Call(parms)

}
func main() {
	var a student = student{
		age:   27,
		name:  "Tom",
		score: 89.5,
	}

	testStruct(a)
}
