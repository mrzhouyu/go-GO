package main

import "fmt"
import "time"

//定义一个channel 用于协程之间通信
var ch = make(chan int)

//定义一个打印机
func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

//使用者1
func Person1() {

	Printer("one world")
	ch <- 666 //调用完给与值 停止阻塞
}

//使用者2
func Person2() {
	<-ch //一开始ch里是没有数据的 那么就会阻塞下面的语句执行 等Person被赋值666 就开始往下执行
	Printer("one dream")
}

//打印机问题代码化
func main() {
	//二者本来都是协程 理论上应该同时都执行 不互相干扰 加入channel后协程之间就可以通信和数据传递了
	go Person1()
	go Person2()

	for { //这里必须加类似于此的语法 让主协程循环执行 相当于阻塞功能 否则主进程直接执行了  子协程就无法执行 后面用channel来实现这种阻塞

	}
}
