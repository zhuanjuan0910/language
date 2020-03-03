package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data1, err := os.Open("./abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	if data1 != nil {
		fmt.Println("fiile open success")
		defer func(file *os.File) { file.Close() }(data1)
	}
	data2, err := ioutil.ReadAll(data1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data2))
	data3, err := ioutil.ReadFile("./abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data3))
}
