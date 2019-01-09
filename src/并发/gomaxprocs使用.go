package main

import (
	"fmt"
	"runtime"
)

const (
	n = 1
)

//主协程退出 子的也退出  注意：可能导致子协程每调用就退出 main函数里面已经执行到了{}尾部
func main() {

	c := runtime.GOMAXPROCS(n) //指定以单核n=1 n越大 任务之间切换越快
	fmt.Println("C = ", c)
	//一下代码可以测试n对任务切换的直观影响
	for {
		go fmt.Print(1)
		fmt.Print(0)
	}

}
