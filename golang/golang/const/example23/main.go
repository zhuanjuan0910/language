package main

import "fmt"

func valueTest(x int) int {
	return x + 1
}
func ptrTest(y *int) int {
	return *y + 2
}
func main() {
	v := 2
	fmt.Println(valueTest(v))
	//fmt.Println(valueTest(&v)) 普通函数值类型传指针类型会报错
	fmt.Println(ptrTest(&v))
	//fmt.Println(ptrTest(v))普通函数指针类型传值类型也不行
}
