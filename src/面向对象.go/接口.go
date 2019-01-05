package main

import "fmt"

type Humaner interface {
	sayHello() //只负责声明不负责实现方法 由别的自定义类型实现
}

type Student struct {
	name string
	age  int
}

func (s *Student) sayHello() {
	fmt.Println("Student被调用")
}

type Teacher struct {
	group string
}

func (t *Teacher) sayHello() {
	fmt.Println("teacher被调用")
}

type Mystr string

func (m *Mystr) sayHello() {
	fmt.Println("mystr被调用了")
}

//定义一个普通函数函数的类型为接口类型
//只有一个函数可以实现不同的表现、多态
func WhoSayHello(i Humaner) {
	i.sayHello()
}

func main() {
	//定义接口类型
	var i Humaner

	//只要实现了此接口的方法类型，那么这个类型的变量(接收者)就可以给i赋值
	s := &Student{"jimi", 22}
	i = s
	i.sayHello()
	//利用多态实现
	WhoSayHello(s)

	t := &Teacher{"bj"}
	i = t
	t.sayHello()
	//利用多态实现
	WhoSayHello(t)
	var str Mystr = "hello world"
	i = &str
	i.sayHello()
	//利用多态实现
	WhoSayHello(&str)

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str
	for _, h := range x {
		h.sayHello()
	}

}
