package main

import (
	"fmt"
	"os"
)

func main() {
	file1, err := os.Open("./abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	if file1 != nil {
		fmt.Println("open file is success")
		defer func(file *os.File) { file.Close() }(file1)
	}
	b1 := make([]byte, 102)
	space1, err := file1.Read(b1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("file read success , 读取 %d 字节。\n", space1)
	fmt.Printf("读取内容：\n%s\n", string(b1))

	b2 := make([]byte, 205)
	space2, err := file1.ReadAt(b2, int64(space1))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("file readat success , 读取 %d 字节。\n", space2)
	fmt.Printf("读取内容：\n%s\n", string(b2))
}
