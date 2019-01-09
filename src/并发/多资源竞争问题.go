package main

import "fmt"
import "time"

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
}

//使用者2
func Person2() {
	Printer("one dream")
}

//打印机问题代码化
func main() {
	//两个函数都是调用Printer 会出现资源竞争  打印出如下字符:    oonnee  wdrorealdm
	go Person1()
	go Person2()
	for { //这里必须加类似于此的语法  来让主进程不执行完毕 否则主进程直接执行了  子协程就无法执行

	}
}
