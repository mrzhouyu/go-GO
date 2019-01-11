package main

import "fmt"
import "time"

func main() {
	timer := time.NewTimer(time.Second) //设置1s
	timer.Reset(3 * time.Second)        //重新设置为3s

	go func() {
		<-timer.C //管道阻塞  时间到才可以继续执行
		fmt.Println("定时器时间到 可以打印")

	}()
	timer.Stop() //停止定时器 子协程不被执行

	for {

	}

}
