package main

import "fmt"

type T struct {
	int
}

func (t T) test() {
	fmt.Println("类型 T 方法集包含全部 receiver T 方法。")
}
func main() {
	u := T{1}
	fmt.Printf("u is %v\n", u)
	u.test()
}
