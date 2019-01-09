package main

import (
	"fmt"
	"time"
)

//有缓存的channel 是放满channel以后阻塞当前程序 直到被被其他协程读取

func main() {
	ch := make(chan int, 3) //定义channel容量为3
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch))

	//新建子协程
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i //三个时候 阻塞等待2schannel被读取 继续执行
			fmt.Println("子协程i= ", i)
		}
	}()
	//故意sleep 错开第一次协程的读取时间
	time.Sleep(time.Second * 2)

	for i := 0; i < 10; i++ {
		res := <-ch //管道内没有内容阻塞当前fou循环 这里的取跟 匿名函数里的赋值是同时的
		fmt.Printf("i=%d时候收到的值是：%d\n", i, res)
	}

}
