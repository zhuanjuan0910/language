package main

import "fmt"

type user struct {
	name string
	age  int
}
type test struct {
	user //匿名字段 说明test结构体有user类的方法。可以直接调用。通过匿名字段可以获得和继承类似的复用能力
}

func (u *user) people() {
	fmt.Println(u.name, u.age)
}
func main() {
	a := test{user{"zhangsan", 27}}
	a.people()
}
