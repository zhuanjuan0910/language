package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func testMap() (ret string, err error) {
	var m map[string]interface{}
	m = make(map[string]interface{})
	m["name"] = "hfvj"
	m["age"] = 13
	m["id"] = 26
	data, err := json.Marshal(m)
	if err != nil {
		fmt.Println("json.Marshal is failed", err)
		return
	}
	ret = string(data)
	return

}
func test() {
	data, err := testMap()
	if err != nil {
		fmt.Println("test is failed", err)
		return
	}
	var m map[string]interface{}
	err = json.Unmarshal([]byte(data), &m) //json反序列化为map，Unmarshal需要传指针
	if err != nil {
		fmt.Println("json.Unmarshal is failed", err)
		return
	}
	fmt.Println(m)
	fmt.Println(reflect.TypeOf(m))
}
func main() {
	test()
}
