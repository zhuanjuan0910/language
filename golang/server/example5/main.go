package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://www.baidu.com/") //get 得到*response
	if err != nil {
		fmt.Println("链接失败", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("读取失败", err)
		return
	}
	fmt.Println(string(data))
}
