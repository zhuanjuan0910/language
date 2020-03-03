package main

import "fmt"

type user struct {
	id   int
	name string
}

func (self *user) people() {
	fmt.Printf("%p,%v\n", self, self)
}
func main() {
	u1 := user{1, "zhangsan"}
	u1.people()
	u2 := u1.people //把一个函数值 赋值给u2变量，绑定实参，显式传递
	u2()
	u3 := (*user).people
	u3(&u1)
}
