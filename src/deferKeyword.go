package main

import "fmt"

func breakRAM(x int) {
	result := 1 / x
	println(result)

}

//延迟函数
//延迟调用函数defer 在函数结束前瞬间才调用
//多个defer遵循"先进后调用" 并且 defer会忽略错误 即使出错 其他的defer也会继续调用
func main() {
	defer fmt.Println("defer one")
	fmt.Println("no defer")
	defer println("defer two")
	//去掉defer执行该函数崩溃后后面的不执行
	defer breakRAM(0)
	defer fmt.Println("defer three")
}
