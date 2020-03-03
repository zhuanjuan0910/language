package main

import (
	"fmt"
	"net/http"
	"os"
	"text/template"
)

type Person struct {
	Name string
	Age  string
}

var myTemplate *template.Template

func initTemplate(fileName string) (err error) {
	myTemplate, err = template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("failed:", err)
		return
	}
	return
}
func userInfo(w http.ResponseWriter, r *http.Request) {
	p := Person{"Tom", "25"}

	myTemplate.Execute(os.Stdout, p)
	file, err := os.OpenFile("./abc.go", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("file open failed:", err)
		return
	}
	myTemplate.Execute(file, p)
}
func main() {
	initTemplate("../example1/index.html")
	http.HandleFunc("/user/info", userInfo)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

//渲染存储文件
