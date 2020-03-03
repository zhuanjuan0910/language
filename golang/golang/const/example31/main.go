package main

import "fmt"

type I interface {
	method1() int
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
	var s I = user{"zhangshan", 20} //只要这个变量含有接口中的方法，该变量就实现了这个接口
	age := s.method1()
	name := s.method2()
	fmt.Println(age, name)
}
