package main

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle hello")
	fmt.Fprintf(w, "----- hello ")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("handle login")
	fmt.Fprintf(w, "----- login ")
}

func main() {
	http.HandleFunc("/", Hello)
	http.HandleFunc("/user/login", login)
	err := http.ListenAndServe("0.0.0.0:8880", nil)
	if err != nil {
		fmt.Println("http listen failed")
	}
}

// 1.运行该代码
// 2.在浏览器输入：http://localhost:8880/user/login或http://localhost:8880
