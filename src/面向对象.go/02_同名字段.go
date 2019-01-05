package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型没有名字 匿名字段
	class  string
	id     int
	name   string
}

func main() {
	var s Student

	s.name = "mike"
	s.id = 100
	fmt.Printf("s == %+v\n", s) //同名的name遵循就近原则 先赋值的是Student里的
	//现式调用 如果非要调用赋值Person里的name
	s.Person.name = "Tom"
	fmt.Printf("s == %+v\n", s)

}
