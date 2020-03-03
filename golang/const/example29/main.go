package main

import "fmt"

type user struct {
	id   int
	name string
}

func (self user) test() {
	fmt.Println(self)
}
func main() {
	u1 := user{2, "tom"}
	myValue := u1.test //立即复制receiver,不是指针类型，不受后续修改影响
	u1.id, u1.name = 2, "jhon"
	u1.test()
	myValue()

}
