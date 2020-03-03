package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"username"`
	Age  int
	Id   int
}

func TestStruct() {
	User1 := User{"Tom", 16, 28}
	data, err := json.Marshal(User1) //结构体转json
	if err != nil {
		fmt.Println("json.Marsha is failed")
		return
	}
	fmt.Printf("%s\n", string(data))
}
func main() {
	TestStruct()
}
