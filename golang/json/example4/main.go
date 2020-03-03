package main

import (
	"encoding/json"
	"fmt"
)

func testSlice() {
	var m map[string]interface{}
	var s []map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "Tom"
	m["age"] = 26
	m["id"] = 15
	s = append(s, m)
	m = make(map[string]interface{})
	m["name"] = "john"
	m["age"] = 24
	m["id"] = 24
	s = append(s, m)
	data, err := json.Marshal(s) //slice转json
	if err != nil {
		fmt.Println("json.Marshal is failed", err)
		return
	}
	fmt.Printf("%s\n", string(data))
}
func main() {
	testSlice()
}
