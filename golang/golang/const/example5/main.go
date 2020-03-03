package main

import "fmt"

const ( //在同一行中的iota值是一样的
	a, b = iota, iota << 10 //0,0
	c, d                    //1,1024
)

func main() {
	fmt.Println(a, b, c, d)
}
