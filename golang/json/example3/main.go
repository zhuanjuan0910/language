package main

import (
	"encoding/json"
	"fmt"
)

func testInt() {
	var a int = 100
	data, err := json.Marshal(a) //int转json
	if err != nil {
		fmt.Println("json.Marshal is failed ", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}
func main() {
	testInt()
}
