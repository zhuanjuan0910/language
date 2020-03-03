package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		fmt.Println(string(buf[0:n]))

	}
}

func main() {
	fmt.Println("start server----")
	listen, err := net.Listen("tcp", "0.0.0.0:50000")
	if err != nil {
		fmt.Println("listen is failed", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("连接失败", err)
			continue
		}
		go process(conn)
	}
}

//服务端
// a. 监听端口:listen,err:=net.Listen("协议"，"端口号"")
// b. 接收客户端的链接:conn,err:=listen.Accept()
// c. 创建goroutine，处理该链接
