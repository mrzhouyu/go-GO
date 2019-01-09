package main

func main() {
	//定义一个双向的channel
	ch := make(chan int) //或者var ch chan
	//可以隐式的转化为单向的channel
	var writeCh chan<- int = ch //只能写不能读
	var readCh <-chan int = ch  //只能读不能写
	writeCh <- 666
	<-readCh

	//单向无法转化为双向 下面语句是错误的
	var ch2 chan int = writeCh
}
