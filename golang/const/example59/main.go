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

func (s student) Print() {
	fmt.Println(s)
}
func testStruct(a interface{}) {
	val := reflect.ValueOf(a)
	kd := val.Kind()
	fmt.Println(val, kd)
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("val is not struct")
		return
	}
	fleds := val.Elem().NumField()
	fmt.Println("struct has ", fleds)
	for i := 0; i < fleds; i++ {
		fmt.Printf("%d%v\n", i, val.Elem().Field(i).Kind())
	}
	methods := val.NumMethod()
	fmt.Printf("struct has %d method", methods)
	var parms []reflect.Value
	val.Elem().Method(0).Call(parms)
}
func main() {
	var a student = student{
		name:  "Tom",
		age:   17,
		score: 89.6,
	}
	testStruct(&a)
}
