package main

import (
	"fmt"
	"reflect"
)

func main() {
	type User struct {
		Name string `json:"user_name" name:"user name"`
		age  int
	}
	u := User{
		Name: "Murphy",
		age:  15,
	}
	f := reflect.TypeOf(u).Field(0)
	fmt.Println(f.Tag.Get("json"))
	fmt.Println(f.Tag.Get("name"))
}
