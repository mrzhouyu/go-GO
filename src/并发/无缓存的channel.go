package main

import (
	"fmt"
	"time"
)

//无缓存的channel 特点是 放取是同时的 channel本身不缓存临时的任何东西  len(ch) cap(ch)都为0  channel没有被放入的时候会阻塞下面的语句
// 当channel被赋值后  channel同时间没有被其他的协程读取时候 也会阻塞当前协程后面的语句

func main() {
	ch := make(chan int) //不指定容量或者指定容量为0 make(chan int,0) 是无缓存的channel len(ch) cap(ch)都为0
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch))

	//新建子协程
	go func() {
		for i := 1; i < 4; i++ {
			fmt.Println("子协程i= ", i)
			ch <- i                 //给与值得同时需要有其他的协程取值 若没有被取 阻塞当前for循环
			time.Sleep(time.Second) //这里加个sleep 可以让打印顺序进行 不然因为程序执行很快 打印信息会比较乱  也可以不加
		}
	}()
	//故意sleep 错开第一次协程的读取时间
	time.Sleep(time.Second * 2)

	for i := 1; i < 4; i++ {
		res := <-ch //管道内没有内容阻塞当前fou循环 这里的取跟 匿名函数里的赋值是同时的
		fmt.Printf("i=%d时候收到的值是：%d\n", i, res)
	}

}
