package main

import "fmt"

func main() {

	a := 1
	b := "one"
	//第一种调用
	f := func() {
		fmt.Printf("a=%d,b=%s\n", a, b)
	}
	f()
	//第二种调用方法
	//定义函数类型 无参数无返回值
	type funType func()
	var f1 funType
	f1 = func() {
		fmt.Printf("a=%d,b=%s\n", a, b)
	}
	f1()
	//第三种调用方式
	func() {
		fmt.Printf("a=%d,b=%s\n", a, b)
	}()
	//带参数无返回
	f4 := func(x, y int) {
		fmt.Printf("a=%d,b=%d\n", x, y)
	}
	f4(3, 4)
	//带参数 无返回
	func(x, y int) {
		fmt.Printf("a=%d,b=%d\n", x, y)

	}(10, 20)
	//单参数且返回
	s, t := func(x, y int) (c, d int) {
		c = x
		d = y
		return
	}(100, 200)
	fmt.Printf("s=%d,t=%d\n", s, t)
}
