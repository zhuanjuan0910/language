package main

import "fmt"

func test() (int, int) {
	return 1, 2
}
func sum(n ...int) int {
	var x int
	for _, i := range n {
		x += i
	}
	return x
}
func add(x, y int) int {
	return x + y
}
func main() {
	fmt.Println(sum(test())) //可将多返回值的函数当做其它函数的实参用
	fmt.Println(add(test()))
}
