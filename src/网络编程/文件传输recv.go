package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func ReceiveData(name string, conn net.Conn) {
	f, err := os.Create(name)
	if err != nil {
		fmt.Println("os.Create err1:", err)
		return
	}

	data := make([]byte, 1024*4)
	//这里for循环表示一直收 客户端同样也有个一直发送的函数
	for {
		n, err1 := conn.Read(data)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件传输结束")
			} else {
				fmt.Println("conn.Read err1:", err1)
			}
			return
		}
		if n == 0 {
			fmt.Println("文件接收完毕")
			break
		}
		f.Write(data[:n])

	}

}

func main() {
	//监听客户端请求
	listener, err1 := net.Listen("tcp", "127.0.0.1:777")
	if err1 != nil {
		fmt.Println("net.Listener err1:", err1)
		return
	}

	defer listener.Close()
	//阻塞等待用户连接 直到连接返回给conn
	conn, err2 := listener.Accept()
	if err2 != nil {
		fmt.Println("listener.Accept err2:", err2)
		return
	}

	defer conn.Close()

	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.Read err3:", err3)
		return
	}
	//接收到文件名字 返回给请求端可以发送数据内容
	name := string(buf[:n]) //文件名
	//回复客户端
	conn.Write([]byte("ok"))
	//发送ok后执行以下函数
	ReceiveData(name, conn) //文件建立与写入

}
