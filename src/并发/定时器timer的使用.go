package main

import "fmt"
import "time"

func main() {
	//新建一个定时器 时间到了只会响应一次  设置2s
	timer := time.NewTimer(time.Second * 2)
	fmt.Println("当前时间：", time.Now())

	//2s后往time.C的管道写入数据，有数据就可以读取 否则阻塞
	t := <-timer.C
	fmt.Println("t= ", t)
}
