package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	quit := make(chan bool)
	//子协程
	go func() {
		//for循环 一直检测
		for {
			select {
			case flag := <-ch:
				fmt.Println("flag=", flag)
			case <-time.NewTimer(time.Second * 4).C: //当定时时间到 此处不再阻塞
				fmt.Println("超时")
				quit <- true //给与值
			}

		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit //定时器没到之前 一直阻塞 主函数不会退出 直到quit有值 主函数执行到  "}" 处
	fmt.Println("超时间退出")
}
