package main

import (
	"fmt"
	"html/template"

	// "io"
	"net/http"
)

type Person struct {
	Name string //需要导出，必须大写
	Age  string
}

var myTemplate *template.Template

func initTemplate(fileName string) (err error) {
	myTemplate, err = template.ParseFiles(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
func userInfo(w http.ResponseWriter, r *http.Request) { //要是指针类型的*http.Request
	p := Person{"Tom", "15"}
	myTemplate.Execute(w, p)
}
func main() {
	initTemplate("../example1/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

//渲染浏览器
