package main

import "fmt"

//程序直接奔溃并且提示错误原因  等级比error高   程序内部会自己实现一个 也可以自定义

func A() {
	fmt.Println("AAA")

}

func B() {
	fmt.Println("BBB")
	panic("调用了panic错误处理")
}

func main() {
	B()
	A()
}
