package main

import (
	"io"
	// "log"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
                    <input type="text" name="in"/>
                    <input type="text" name="in"/>
                     <input type="submit" value="Submit"/>
             </form></body></html>`

func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		request.ParseForm()
		io.WriteString(w, request.Form["in"][1])
		io.WriteString(w, "\n")
		io.WriteString(w, request.FormValue("in"))
	}
}

/*
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
*/
func main() {
	http.HandleFunc("/test2", FormServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
	}
}
