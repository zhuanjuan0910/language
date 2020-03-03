package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("1", "one")  //增
	v, ok := m.Load("1") //查
	fmt.Println(m.Load("1"))
	m.Store("1", "oneone") //改
	fmt.Println(m.Load("1"))
	v, ok = m.LoadOrStore("2", "Two") //查，没有就增
	fmt.Println(v, ok)
	m.Range(func(k, v interface{}) bool { //遍历
		fmt.Println(k, v)
		return true
	})
	m.Delete("1")                         //删除
	m.Range(func(k, v interface{}) bool { //遍历
		fmt.Println(k, v)
		return true
	})

}
