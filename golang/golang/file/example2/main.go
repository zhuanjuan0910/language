package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	file1, err := os.Create("./file1.go")
	if err != nil {
		fmt.Println("create file is", err)
	}
	if file1 != nil {
		defer func(file *os.File) { file.Close() }(file1)
		fmt.Println("file1 create success")
	}
	filename := file1.Name()
	fmt.Printf("filename is %v\n", filename)
	file2 := os.NewFile(uintptr(syscall.Stdin), "./file2.go")
	if file2 != nil {
		defer func(file *os.File) { file.Close() }(file2)
		fmt.Println("file2 create success")
	}
	fileInfo1, err := file1.Stat()
	if err != nil {
		fmt.Printf("get file1 info err : %v\n", err)
	}
	fmt.Println(fileInfo1)

	fileInfo2, err := file2.Stat()
	if err != nil {
		fmt.Printf("get file2 info err : %v\n", err)
	}
	fmt.Println(fileInfo2)
	b := os.SameFile(fileInfo1, fileInfo1)
	if b {
		fmt.Println("fileInfo1 与 fileInfo1 是同一个文件")
	} else {
		fmt.Println("fileInfo1 与 fileInfo1 不是同一个文件")
	}

	fileMode1 := fileInfo1.Mode()
	b = fileMode1.IsRegular()
	if b {
		fmt.Println("file1 是普通文件")
	} else {
		fmt.Println("file1 不是普通文件")
	}

	b = fileMode1.IsDir()
	if b {
		fmt.Println("file1 是普通目录")
	} else {
		fmt.Println("file1 不是普通目录")
	}
}
