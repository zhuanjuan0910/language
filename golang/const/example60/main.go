package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"stu_name"`
	Age   int
	Score float32
}

func testStruct(a interface{}) {
	val := reflect.TypeOf(a)
	fmt.Println(val)
	fmt.Println(val.Elem().Field(0))
	tag := val.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag:%s\n", tag)
}
func main() {
	var v Student = Student{
		Name:  "Tom",
		Age:   24,
		Score: 89.6,
	}
	testStruct(&v)
}
