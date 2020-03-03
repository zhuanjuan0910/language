package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	str := "hello world, hello golang"
	err := ioutil.WriteFile("./abc.go", []byte(str), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")
}
