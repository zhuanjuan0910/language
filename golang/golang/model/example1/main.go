package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	age  string
}

func main() {
	t, err := template.ParseFiles("./index.html") //得到一个模板指针对象
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	p := Person{Name: "Mary", age: "31"}
	if err := t.Execute(os.Stdout, p); err != nil { //exeecute方法把p填充到t，并写入到标准输出
		fmt.Println("There was an error:", err.Error())
	}

}

//渲染终端
