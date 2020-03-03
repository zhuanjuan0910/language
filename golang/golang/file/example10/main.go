package main

import (
	"archive/tar"
	"fmt"
	"os"
)

func main() {
	fileName := "./abc.tar.gz"
	intersetByte := []byte("this is a test tar")
	file1, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	if file1 != nil {
		defer func(file *os.File) { file.Close() }(file1)
	}
	data := tar.NewWriter(file1)
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		fmt.Printf("open file ./file.tar.gz err : %v\n", err)
		return
	}
	head, err := tar.FileInfoHeader(fileInfo, "")
	if err != nil {
		fmt.Printf("write WriteHeader err : %v\n", err)
		return
	}
	head.Size = int64(len(intersetByte))
	err = data.WriteHeader(head)
	if err != nil {
		fmt.Printf("write WriteHeader err : %v\n", err)
		return
	}
	ret, err := data.Write(intersetByte)
	if err != nil {
		fmt.Printf("write ./file.tar.gz err : %v\n", err)
		return
	}
	fmt.Println(ret)

	err = data.Flush()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = data.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
