package main

import (
	"fmt"
	"runtime"
)

func test() {
	fmt.Println("three...............................")
	//return //return表示到此结束不打印four
	//换成goExit表示终止整个协程
	runtime.Goexit()
	fmt.Println("four.......................................")
}

//主协程退出 子的也退出  注意：可能导致子协程每调用就退出 main函数里面已经执行到了{}尾部
func main() {

	go func() {
		fmt.Println("one。。。。。。。。。。。。。。。")
		test() //这里调用函数 包括goExit 那么two也不会执行 因为是终止整个协程
		fmt.Println("two...................")
	}()

	for { //故意写一个无线循环

	}
}
