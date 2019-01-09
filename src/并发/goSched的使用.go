package main

import (
	"fmt"
	"runtime"
)

//主协程退出 子的也退出  注意：可能导致子协程每调用就退出 main函数里面已经执行到了{}尾部
func main() {

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子")
		}
	}()

	for i := 1; i < 5; i++ { //如果没有加定时 主程序执行执行完退出 子程序来不及执行 所有这里使用gosched让出CPU
		runtime.Gosched() //让出CPU让其他的任务先执行
		fmt.Println("Master")
	}

}
