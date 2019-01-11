package main

import "net"
import "fmt"
import "time"

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:1000") //服务器监听IP端口
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	defer listener.Close() //函数关闭前关闭监听
	//for循环等待(阻塞)客户链接
	for {
		fmt.Println("waiting....")
		time.Sleep(time.Second)
		conn, err1 := listener.Accept() //等待来自监听的数据 没有数据在此阻塞
		if err1 != nil {
			fmt.Println("err1= ", err1)
			continue
		}
		//接收用户请求
		fmt.Println("监听到来自客户端的请求...")
		buf := make([]byte, 1024) //定义缓存区 1024 byte型切片
		n, err2 := conn.Read(buf) //读取数据 返回真实数据长度和err
		if err2 != nil {
			fmt.Println("err2= ", err2)
			continue
		}
		fmt.Println("buf为：", string(buf[:n]))

	}

}
