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
	var Y Humaner //子集
	var y Yuchou  //超集
	j := &Student{"yuchou", 24}
	y = j
	Y = y //超集转换为子集
	Y.sayHello()
}
