package main

import "fmt"

type User struct {
	id   int
	name string
}

func main() {
	u := User{1, "Tom"}
	var vi, pi interface{} = u, &u

	// vi.(User).name = "Jack"       // Error: cannot assign to vi.(User).name
	pi.(*User).name = "Jack" //接口类型只是复制对象，只有指针才能改变其状态

	fmt.Printf("%v\n", vi.(User))
	fmt.Printf("%v\n", pi.(*User))
}
