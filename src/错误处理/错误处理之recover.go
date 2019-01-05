package main

import "fmt"

func funA() {
	fmt.Println("function A")
}
func funB(l int) {
	//defer放哪里都一样 函数调用结束才执行
	//这里的匿名函数 定义且调用
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var arr [10]int
	arr[l] = 100

}
func funC() {
	fmt.Println("function C")
}

func main() {
	funA()
	funB(15)
	funC()
}
