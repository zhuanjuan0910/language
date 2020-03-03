package main

import (
	"fmt"
	"sync"
)

var m sync.Map

type userInfo struct {
	Name string
	Age  int
}

func main() {
	v, ok := m.LoadOrStore("1", "one")
	fmt.Println(v, ok)

	v, ok = m.Load("1")
	fmt.Println(v, ok)
	v, ok = m.LoadOrStore("1", "oneone")
	fmt.Println(v, ok)
	m.Store("1", "oneone")
	fmt.Println(m.Load("1"))
	m.Store("2", "Two")
	fmt.Println(m.Load("2"))
	m.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
	map1 := make(map[string]userInfo)
	var user1 userInfo
	user1.Name = "ChamPly"
	user1.Age = 24
	map1["user1"] = user1

	var user2 userInfo
	user2.Name = "Tom"
	user2.Age = 18
	m.Store("map_test", map1)

	mapValue, _ := m.Load("map_test")

	for k, v := range mapValue.(interface{}).(map[string]userInfo) {
		fmt.Println(k, v)
		fmt.Println("name:", v.Name)
	}
}
