package main

import (
	"fmt"
	"os"
)

func main() {
	file1, err := os.Open("./file1.go")
	if err != nil {
		fmt.Println("open file1 is ", err)
	}
	if file1 != nil {
		defer func(file *os.File) { file.Close() }(file1)
	}
	file2, err := os.OpenFile("./file2.go", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file2 is ", err)
	}
	if file2 != nil {
		defer func(file *os.File) { file.Close() }(file2)
	}
	b1 := []byte("hello world")
	off, err := file2.Write(b1)
	if err != nil {
		fmt.Printf("file2 write err : %v\n", err)
	}
	fmt.Printf("file2 write success , off is %v\n", off)
	b2 := []byte("hello golang")
	off, err = file2.WriteAt(b2, int64(off))
	if err != nil {
		fmt.Printf("file2 writeat err : %v\n", err)
	}
	fmt.Printf("file2 writeat success , off is %v\n", off)

	str := "this is write string test .\n"
	off, err = file2.WriteString(str)
	if err != nil {
		fmt.Printf("file2 write string err : %v\n", err)
	}
	fmt.Printf("file2 write string success , off is %v\n", off)

}
