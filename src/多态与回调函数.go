package main

import "fmt"

type FuncType func(int, int) int

func Add(a int, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}
func Mult(a, b int) int {
	return a * b
}

func Calc(a int, b int, testFun FuncType) (result int) {
	result = testFun(a, b)
	return
}

func main() {
	var result int
	//利用多态这里可任意传入加减乘除函数 Calc(2,1,Minud)  Calc(2,1,Mult)
	//调用同一个接口实现多种
	result = Calc(2, 1, Add)
	fmt.Println("result = ", result)
}
