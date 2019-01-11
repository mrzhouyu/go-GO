package main

import "fmt"

//这里的ch值写  exit只读 双向通道转单向
func fibonacci(ch chan<- int, exit <-chan bool) {
	x, y := 1, 1

	for {
		//监控channel数据流
		select {
		case ch <- x: //此处会执行8次后阻塞 因为主函数子协程 只接收8次
			x, y = y, x+y
		case <-exit: //此处阻塞 直到 exit被主函数子协程赋值
			fmt.Println("收到退出信号，退出进程")
			return //终止整个函数
		}
	}

}
func main() {
	ch := make(chan int)    //无缓存channel 用于产生数据 有产生必须有接收
	exit := make(chan bool) //be used to exit

	//新建协程消费者 读取ch内容
	go func() {
		for i := 0; i < 8; i++ {
			num := <-ch
			fmt.Println("num=", num)
		}
		exit <- true //读取8次后 给exit赋值

	}()
	//生产者
	fibonacci(ch, exit) //传递方式的相当于地址传递 函数里的值会影响外部的值
}
