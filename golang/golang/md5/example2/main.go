package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "hello world,hello golang"
	fmt.Println(str)
	s := []byte(str)
	fmt.Printf("byte is %v\n", s)
	encode := base64.StdEncoding.EncodeToString(s) //base64编码
	fmt.Printf("encode base64 %v\n", encode)
	decode, err := base64.StdEncoding.DecodeString(encode) //base64解码
	if err != nil {
		fmt.Println("解码失败", err)
		return
	}
	fmt.Println(decode)
	//如果要用在url中，需要使用URLEncoding
	urlEncode := base64.URLEncoding.EncodeToString(s)
	fmt.Printf("encode base64 %v\n", encode)
	urlDecode, err := base64.URLEncoding.DecodeString(urlEncode)
	if err != nil {
		fmt.Println("解码失败", err)
		return
	}
	fmt.Println(urlDecode)

}
