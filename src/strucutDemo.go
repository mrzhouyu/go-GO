package main

import "fmt"

type myStr struct {
	name   string
	age    int
	height int
}

var ins1 myStr

func main() {
	//复制方式1
	ins1.age = 18
	ins1.height = 174
	ins1.name = "zhouyu"
	fmt.Println("out my strcut:", ins1)
	//赋值方式2 调用结构体方法1
	ins2 := myStr{
		name:   "chenyajinh",
		age:    21,
		height: 158,
	}
	fmt.Println("outs my strcut: ", ins2)
	ins2.test()
	//调用结构体方法指针
	ins3 := &myStr{
		name:   "chenyajinh",
		age:    21,
		height: 158,
	}
	fmt.Println("outs my strcut: ", ins3)
	ins3.testPtr()
}

//结构体得方法1 就是给该结构体添加一个方法 下同
func (this myStr) test() {
	fmt.Println("function test")
	return
}

//结构体得方法2 指针方式
func (this *myStr) testPtr() {
	fmt.Println("functions testPtr")
	return
}
