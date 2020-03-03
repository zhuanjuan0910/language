package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("../example8/abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("删除成功")
	err = os.RemoveAll("../example7/abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("也删除成功")

}
