package main

import (
	"fmt"
)

//所有变量初始化为0  int ->0   string -> "" bool ->false *int ->nil
//三种定义变量得方法
var a int = 10086

var b = 10000

// c := 00001 //改定义方法必须在函数内部 只能用来定义局部变量

func main() {
	c := 100001
	fmt.Print("my some value:", a)
	fmt.Print("second value :", b)
	fmt.Print("third value:", c)
}
