package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "149.248.3.74:1111")
	if err != nil {
		fmt.Println("net.Dial err: ", err)
		return
	}

	//主协程负责连接相关 这样只要断开连接 子协程也会退出
	//获取键盘输入的数据
	go func() {
		str := make([]byte, 1024)
		for {
			n, err2 := os.Stdin.Read(str) //从键盘读取数据
			if err2 != nil {
				fmt.Println("os.Stdin.Read err2 :", err2)
				return
			}
			// fmt.Printf("发送：%s", string(str[:n]))
			conn.Write(str[:n-1])
		}
	}()

	//主协程 接收服务器返回的数据

	//切片缓冲区
	buf := make([]byte, 1024)
	//不停地等待来自服务器返回的数据
	for {
		n, err1 := conn.Read(buf)
		if err1 != nil {
			fmt.Println("conn.Read err1: ", err1)
			return
		}
		fmt.Printf("来自服务器的消息：%s", string(buf[:n]))
	}
	defer conn.Close()
}
