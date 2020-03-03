package main

type Color int

const (
	Black Color = iota
	Red
	Blue
)

func test(c Color) {}

func main() {
	c := Black
	test(c)

	// x := 1
	// test(x)整型不能传给color类型
	// ./main.go:18:6: cannot use x (type int) as type Color in argument to test

	test(1) // 常量会被编译器自动转换。
}
