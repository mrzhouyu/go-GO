package main

import "fmt"

//生产者 只能写 不能读
func Producer(out chan<- int) {
	for i := 1; i < 10; i++ {
		out <- i
	}
	close(out) //写结束关闭
}

//只能读不能写
func Consumer(in <-chan int) {
	for i := range in {
		fmt.Println("i= ", i)
	}
}

func main() {
	ch := make(chan int) //定义一个无缓存双向通道
	//协程1
	go Producer(ch) //注意这里不能两个都是子协程 不然默认的主协程会立马退出 子协程也会退出
	//协程2
	Consumer(ch) //主协程
}
