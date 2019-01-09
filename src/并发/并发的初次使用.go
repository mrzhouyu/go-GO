package main

import (
	"fmt"
	"time"
)

func Task() {
	for {
		fmt.Println("这是一个子协程")
		time.Sleep(time.Second)
	}
}

//goroutine 协程
func main() {

	go Task() //子任务 加go关键字

	for { //主任务 主 任务必须放到子任务后面
		fmt.Println("这是主协程")
		time.Sleep(time.Second) //延时1s
	}

}
