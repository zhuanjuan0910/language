package main

import (
	"encoding/json"
	"fmt"
)

func TestMap() {
	mp := make(map[string]interface{})
	mp["name"] = "Tom"
	mp["age"] = 15
	mp["id"] = 67

	data, err := json.Marshal(mp) //map转json
	if err != nil {
		fmt.Println("json.Marsha is failed")
		return
	}
	fmt.Printf("%s\n", string(data))
}
func main() {
	TestMap()
}
