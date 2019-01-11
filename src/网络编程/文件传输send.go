package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func SendFile(path string, conn net.Conn) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}

	defer f.Close()

	//定义一个数据切片缓存 读取多少发送多少 直到读取完毕
	buf := make([]byte, 1024*4)
	for {
		n, err1 := f.Read(buf)
		if err1 != nil {
			if err1 == io.EOF {
				fmt.Println("文件发送完毕")
			} else {
				fmt.Println("f.Read err1:", err1)
			}
			return
		}
		conn.Write(buf[:n]) //给服务器发送内容
	}
}

func main() {
	var path string
	fmt.Print("请输入文件地址（绝对路径）：")
	fmt.Scan(&path) //类似于C语言scanf
	//获取文件名 info.Name
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println("os.Stat err:", err)
		return
	}

	//连接服务器
	conn, err1 := net.Dial("tcp", "127.0.0.1:777")
	if err1 != nil {
		fmt.Println("net.Dial err1:", err1)
		return
	}
	//函数结束前关闭连接
	defer conn.Close()
	//原则上是先发送 文件名 给服务器 请求验证
	_, err2 := conn.Write([]byte(info.Name()))
	if err2 != nil {
		fmt.Println("conn.Write err2:", err2)
		return
	}

	//若得到回复是ok 说明服务器端准备好了 可以发送具体文件数据了
	var n int
	buf := make([]byte, 1024)
	n, err3 := conn.Read(buf)
	if err3 != nil {
		fmt.Println("conn.Read err3:", err3)
		return
	}
	//返回的是ok则可以发送数据
	if "ok" == string(buf[:n]) { //收到ok执行以下函数
		SendFile(path, conn)
	}
}
