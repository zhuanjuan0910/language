package main

import (
	"crypto/hmac"
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Hmac(key, data string) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	return hex.EncodeToString(hmac.Sum([]byte("")))
}
func main() {
	fmt.Println(Hmac("key2", "hello"))
}
