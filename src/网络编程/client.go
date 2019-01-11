package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:1111")
	if err != nil {
		fmt.Println("err =", err)
		return
	}

	defer conn.Close() //函数退出前关闭conn
	buf := make([]byte, 1024)
	for {
		n, err1 := conn.Read(buf)
		if err != nil {
			fmt.Println("err1 ", err1)
			return
		}
		fmt.Println("消息来啦：", string(buf[:n-1]))

	}

}
