package main

import "fmt"

type I1 interface {
	method1() int
}
type I2 interface {
	method2() string
}
type user struct {
	name string
	age  int
}

func (u user) method1() int {
	return u.age
}
func (u user) method2() string {
	return u.name
}
func main() {
	var a I1 = user{"zhangsan", 12}
	fmt.Println(a.method1())
	var b I2 = user{"klis", 23}
	fmt.Println(b.method2())
}
