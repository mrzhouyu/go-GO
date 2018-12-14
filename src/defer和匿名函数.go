package main

import "fmt"

//如下 main main1打印值得顺序不一样
//因为单参数且调用的匿名函数 会先传递参数
func main() {
	a := 1
	b := 2
	defer func() {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}() //()表示调用此匿名函数

	a = 11
	b = 22
	fmt.Printf("a=%d,b=%d\n", a, b)
}
func main1() {
	a := 1
	b := 2
	defer func(a, b int) {
		fmt.Printf("a=%d,b=%d\n", a, b)
	}(a, b) //()表示调用此匿名函数

	a = 11
	b = 22
	fmt.Printf("a=%d,b=%d\n", a, b)

}
