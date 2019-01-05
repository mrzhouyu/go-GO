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
	//结构体变量是一个指针变量，它能够调用的方法有哪些，这些方法就是一个集合，简称方法集
	p := &Person{"jimi", 'w', 18} //p为指针变量 也可以p :=&Person{"jimi", 'w', 18}
	p.SetinfoValue()              //既可以调用值函数 先把指针p转换成*p后再调用 其实就是(*p).SetInfoPointer() 这里是因为go内部做的转换，
	p.SetInfoPointer()            //也可以调用指针的方法
}
