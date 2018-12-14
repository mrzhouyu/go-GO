package main

import "fmt"

func Add(a, b int) int {
	return a + b

}
func minus(a, b int) int {
	return a - b
}

//函数也是一种数据类型,通过type可以给函数命名
//如下 FuncType 它是一种函数类型

type FuncType func(int, int) int //没有函数名字 没有{}

func main() {
	result1 := Add(1, 1)
	fmt.Printf("result1 = %d \n", result1)

	//声明一个函数类型的变量
	var test FuncType
	test = Add
	result2 := test(1, 1) //等价与Add(1,1)
	fmt.Println("result2=", result2)
	test = minus
	result3 := test(2, 1) //等价于minus(2,1)
	fmt.Println("result3=", result3)

}
