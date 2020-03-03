package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file1, err := os.Open("./abc.txt")
	if err != nil {
		fmt.Println(err)
	}
	if file1 != nil {
		fmt.Println("open file success")
		defer func(file *os.File) { file.Close() }(file1)
	}
	reader := bufio.NewReader(file1)
	b1 := make([]byte, 102)
	space1, err := reader.Read(b1)

	if err != nil {
		fmt.Printf("read err : %v\n", err)
	}
	fmt.Printf("read success , 读取 %d 字节\n读取的内容：\n%s\n", space1, string(b1))

	var line []byte
	for {
		data, prefix, err := reader.ReadLine()
		if err == io.EOF {
			fmt.Println(err)
			break
		}

		line = append(line, data...)
		if !prefix {
			fmt.Printf("data:%s\n", string(line))
		}
	}
	fmt.Println(string(line))
}
