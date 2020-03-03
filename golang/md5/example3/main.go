package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	//方法一
	str := "abc"

	byte := []byte(str)
	has := sha1.Sum(byte)
	s := fmt.Sprintf("%x\n", has)
	fmt.Println(s)
	//方法二
	a := sha1.New()
	io.WriteString(a, str)
	b := a.Sum(nil) //将a的hash转成[]byte格式
	c := hex.EncodeToString(b)
	fmt.Println(c)
}
