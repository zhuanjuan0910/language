package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	fileName := "../example10/abc.tar.gz"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file is failed:", err)
		return
	}
	if file != nil {
		defer func(file *os.File) { file.Close() }(file)
	}
	read := tar.NewReader(file)
	head, err := read.Next()
	if err != nil {
		fmt.Println(err)
		return
	}
	var getByte = make([]byte, head.Size)
	_, err = read.Read(getByte)
	if err != nil {
		fmt.Printf("read err : %v\n", err)
	}
	fmt.Println(string(getByte))

}
