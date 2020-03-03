package main

import (
	"fmt"
	"io"
	"os"
)

func copyFile(srcName, dstName string) (write int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open file is,", err)
		return
	}
	if src != nil {
		defer func(file *os.File) { file.Close() }(src)

	}
	dst, err := os.OpenFile(dstName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dst.Close()
	return io.Copy(dst, src) //从一个文件拷贝到另一个文件
}
func main() {
	copyFile("./abc.go", "./hello.go")
	fmt.Println("copy done")
}
