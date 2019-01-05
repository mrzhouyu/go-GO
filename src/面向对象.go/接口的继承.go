package main

import "fmt"

type Humaner interface { //子集
	sayHello() //只负责声明不负责实现方法 由别的自定义类型实现
}

type Yuchou interface { //超集
	Humaner //继承接口
	game()
}

type Student struct {
	name string
	age  int
}

func (s *Student) sayHello() {
	fmt.Println("Student被调用")
}

func (s *Student) game() {
	fmt.Println("LOL")
}

func main() {
	var f Yuchou
	j := &Student{"yuchou", 24}
	f = j
	f.sayHello() //继承的方法
	f.game()     //本作作用域的方法
}
