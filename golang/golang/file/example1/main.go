package main

import (
	"fmt"
	"os"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Printf("get path err : %v\n", err)
	}
	fmt.Printf("当前目录 : %v\n", path)

	err = os.Mkdir("./golang", 0777)
	if err != nil {
		fmt.Printf("mkdir golang err : %v\n", err)
		if os.IsExist(err) {
			fmt.Println("文件已存在！")
		}
		if os.IsNotExist(err) {
			fmt.Println("文件不存在！")
		}
	}

	err = os.MkdirAll("./golang/go", 0777)
	if err != nil {
		fmt.Printf("mkdirall err : %v\n", err)
	}

	err = os.Rename("./golang/go", "./golang/gogo")
	if err != nil {
		fmt.Printf("rename err : %v\n", err)
	}
}
