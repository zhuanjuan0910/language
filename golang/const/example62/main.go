package main

import (
	"fmt"
	"reflect"
)

func main() {
	type student struct {
		Name string `user name`
		Age  string `user age`
	}
	var u = student{
		Name: "Tom",
		Age:  "15",
	}
	s := reflect.TypeOf(u).Field(0).Tag
	fmt.Println(s)
}
