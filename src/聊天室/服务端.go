package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

//用来处理全局来自客户端的消息
var ch chan string

//整个群协程之间通信的管道 定义并且初始化
var message = make(chan string)

//用来存储用户相关信息的结构体
type Cli struct {
	Ch   chan string //用户消息的管道 来自于message的消息
	Addr string      //用户地址 IP:PORT
	Name string      //用户名 默认等于addr
}

//存储用户相关信息的map 定义并且初始化 分配空间
var onlineCliMap = make(map[string]Cli)

//消息构造函数
func MakeMsg(clientNew Cli, msg string) (buf string) {
	buf = "[" + clientNew.Addr + "]" + clientNew.Name + ": " + msg
	return
}

//用户查询
func QueryUser(client Cli) {
	str := "\n" + "用户列表：\n"
	for _, u := range onlineCliMap {
		str = str + u.Name + "\n"
	}
	client.Ch <- MakeMsg(client, str[:len(str)])
}

//处理用户连接函数 可以存储连接用户的信息
func HandelCli(conn net.Conn) {
	//退出处理
	quit := make(chan bool)
	hasData := make(chan bool)
	defer conn.Close()
	//获取客户端网络地址
	cliAddr := conn.RemoteAddr().String()
	//构造结构体 必须有三个字段  第一个字段给一个空的管道值是nil 后面两个都赋值为cliAddr
	clientNew := Cli{make(chan string), cliAddr, cliAddr}
	//把这个结构体 添加到用户map
	onlineCliMap[cliAddr] = clientNew
	//新开一个协程 专门用于 发送消息给当前用户
	go WriteMsgToCli(clientNew, conn)
	//当前用户提示
	clientNew.Ch <- MakeMsg(clientNew, "我上线了")
	//消息构造 告诉当前所有用户 xx上线
	message <- MakeMsg(clientNew, "login")
	//增加一个协程处理 管道流动 实现退出和超时处理

	//用户上线就不退出  实现一直交互 此主协程不能退出 以下实现用户聊天消息广播发送
	buf := make([]byte, 2048)
	go func() {
		for {
			n, err := conn.Read(buf) //等待来自客户端的消息
			if n == 0 {              //用户断开或者退出
				quit <- true
				fmt.Println("sssss err", err)
				return
			}
			//来自当前用户的消息广播出去
			sendMsg := string(buf[:n-1]) //去除换行
			//增加一个用户查询功能和改名功能
			if sendMsg == "use" {
				QueryUser(clientNew)
			} else if len(sendMsg) > 8 && sendMsg[:6] == "rename" {
				s := strings.Split(sendMsg, "|")
				clientNew.Name = s[1]
				clientNew.Ch <- MakeMsg(clientNew, "改名成功")
			} else {
				message <- MakeMsg(clientNew, sendMsg)
			}
			hasData <- true //只要有数据来了就写 一个true 边写边给到 select一个信号 表示此客户未超时
		}

	}()

	//主协程用来处理退出登录这样的 整体任务
	for {
		select {
		case <-quit:
			message <- MakeMsg(clientNew, "用户:"+clientNew.Name+"退出")
			delete(onlineCliMap, cliAddr)
			return
		case <-hasData: //1min内这里能通信 就不会进入到下面的 定时器
		case <-time.After(time.Second * 20):
			message <- MakeMsg(clientNew, "timeout logout")
			delete(onlineCliMap, cliAddr)
			return

		}
	}

}

//发送信息的函数 只要管道有消息来就发送
func WriteMsgToCli(c Cli, conn net.Conn) {
	//一直 读取来自管道的数据  也可以用for {m:=<-c.Ch}这种方式实现一直读取
	for m := range c.Ch {
		conn.Write([]byte(m + "\n")) //发送给当前客户端
	}

}

//广播函数 遍历map
func Manager() {

	for {
		msg := <-message                      //没有消息前 阻塞 有的话执行以下 代码 message 来自于处理用户连接那里 用户上线message就会有值
		for _, client := range onlineCliMap { //client是结构体
			client.Ch <- msg //每个用户的消息管道 都发送一个 上线消息或者其他消息

		}
	}
}

func main() {
	//监听客户端请求
	listener, err := net.Listen("tcp", "127.0.0.1:1111")
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	//记得用完关闭
	defer listener.Close()
	//新建一个协程用户上线消息转发  只要有用户上线下线 都遍历全局的用户map分别发送上线信息 广播
	go Manager()

	for {
		conn, err1 := listener.Accept() //等待用户连接 只要有用户连接 就开一个协程去处理
		if err != nil {
			fmt.Println("listener.Accept err1:", err1)
			continue //出错 继续下一次等待阻塞
		}
		//有用户连接 开启一个协程去处理
		go HandelCli(conn)

	}

}
