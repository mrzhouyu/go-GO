package main

import "fmt"

//封装：通过方法实现    继承：通过匿名函数实现  多态：通过结构实现
type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型没有名字 匿名字段
	class  string
	id     int
}

func main() {

	s := Student{Person{"chou", 'm', 16}, "中二", 1}
	fmt.Println(s)
	s.name = "jing"
	fmt.Println(s)
	s.Person = Person{"yu", 'w', 12}
	fmt.Println(s)
}
