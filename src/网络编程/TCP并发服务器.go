package main

import (
	"fmt"
	// "time"
	"net"
	"strings"
)

//这里传入 的conn的数据类型需要去net包里找到
func DealConn(conn net.Conn) {

	//得到当前链接的客户端的Address  IP and Port
	addr := conn.RemoteAddr().String()
	fmt.Println("连接", addr, "成功!")

	buf := make([]byte, 1024) //定义一个buff 用于接收来自TCP的数据
	//for循环 一直读取
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("读取数据失败 err=", err)
			return
		}
		//函数调用完关闭连接
		defer conn.Close()
		data := string(buf[:n])
		fmt.Printf("%s: %s", addr, data)
		if "exit" == string(buf[:n-2]) { //这里 -1的目的是去除 多余的换行符号\n
			fmt.Printf("接收到来自%s的退出请求\n", addr)
			return
		}
		//数据转换为大写再发送给用户字符串翻转 再转换为 byte切片类型发送
		conn.Write([]byte(strings.ToUpper(data)))
	}
}

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:888")
	if err != nil {
		fmt.Println("err=", err)
		return
	}

	defer listener.Close()

	for {
		conn, err1 := listener.Accept() //阻塞直到有数据传入
		if err1 != nil {
			fmt.Println("err1 = ", err1)
			return
		}
		//定义协程 就可以接收不同的客户端请求 一个用户一个新的协程
		go DealConn(conn)
	}

}
