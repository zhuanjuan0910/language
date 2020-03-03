package main

import "fmt"

//结构体
type user struct {
	name string
	age  int
}

//结构体方法
func (u user) people() {
	fmt.Printf("%v:%d\n", u.name, u.age)
}
func main() {
	//值类型调用方法
	u1 := user{"zhangsan", 19}
	u1.people()
	//指针类型调用fangfa
	u2 := &u1
	u2.people()
}
