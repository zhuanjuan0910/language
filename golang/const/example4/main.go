package main

import "fmt"

const (
	_        = iota             //iota=0
	KB int64 = 1 << (10 * iota) //iota=1
	MB                          //2
	GB
	TB
)

func main() {
	fmt.Println(KB, MB, GB, TB)
}
