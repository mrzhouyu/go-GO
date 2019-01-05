package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

func (p Person) SetinfoValue() {
	fmt.Println("SetinfoValue")
}
func (p *Person) SetInfoPointer() {
	fmt.Println("SetInfoPointer")
}

func main() {
	P := Person{"jimi", 18, 'm'}
	p.SetInfoPointer() //传统调用方法

	//保存方法入口地址
	pFunc := p.SetInfoPointer //这个就是方法值 调用函数时候无需再传递接收者 隐藏了接收者
	pFunc()                   //等价于p.SetInfoPointer()

	//保存方法入口地址
	f := (*Person).SetInfoPointer() //这个是方法表达式
	f(&p)                           //显式把接收者传递过去 相当于p.SetInfoPointer()

	f1 := (Person).SetinfoValue()
	f1.(p)

}
