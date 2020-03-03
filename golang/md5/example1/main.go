package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	//方法一：将字符串转换成md5加密后的十六进制表示法
	var str string = "abc"
	data := []byte(str)
	has := md5.Sum(data)
	md51 := fmt.Sprintf("%x", has) //将字节数组转换为16 进制
	fmt.Println(md51)
	//方法二
	w := md5.New()
	io.WriteString(w, str) //将str写入到w中
	bw := w.Sum(nil)       //w.Sum(nil)将w的hash转成[]byte格式

	// md5str2 := fmt.Sprintf("%x", bw)	//将 bw 转成字符串
	md5str2 := hex.EncodeToString(bw) //将 bw 转成字符串
	fmt.Println(md5str2)

}
