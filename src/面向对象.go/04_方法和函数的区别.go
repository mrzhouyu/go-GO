package main

import "fmt"

//面向过程，实现两个数字相加函数
func Add01(a, b int) int {
	return a + b
}

//面向对象  给某个类型绑定一个函数  (这里给int型绑定一个方法)实现两个数相加的方法
type Int int

func (tmp Int) Add02(data Int) Int {
	return tmp + data
}

//结构体类型添加方法
type Student struct {
	age  int
	name string
	id   int
}

//该方法只有打印功能 普通调用
func (tmp Student) PringInfo() {
	fmt.Println("tmp==", tmp)
}

//修改功能 用到指针

func (p *Student) SetInfo(a int, n string, i int) {
	p.age = a
	p.id = i
	p.name = n
}

//注意 不能绑定原本就是指针类型的量
//如下 为错误的调用方法:
// // type Point *int
// // func (tmp Point) Test(){

// }

func main() {
	//函数的调用
	result := Add01(4, 5)
	fmt.Println(result)

	//方法的调用
	var a Int = 10
	b := a.Add02(10)
	fmt.Println("这是来自于方法得出的结果：", b)
	//本质上实现的功能是一样 只是换了一种表现形式
	//调用结构体的方法
	S := Student{16, "Tom", 1}
	S.PringInfo()
	//修改内容的调用
	var P Student
	(&P).SetInfo(17, "jako", 2) //调用修改方法
	P.PringInfo()               //打印

}

//只有接收者类型不一样 即使方法同名那也是不同的方法
