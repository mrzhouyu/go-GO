package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型没有名字 结构体匿名字段
	int    //基本类型的匿名函数
	class  string
	id     int
}
type Student1 struct {
	*Person //只有类型没有名字 结构体匿名字段
	int     //基本类型的匿名函数
	class   string
	id      int
}

func main() {

	S1 := Student1{&Person{"yu", 'm', 12}, 123, "3", 100} //指针类型的初始化 也可以通过New方法初始化S=New(Student) 就可以调用了

	S := Student{Person{"yu", 'm', 12}, 123, "3", 100} //普通初始化
	fmt.Println(S)
	fmt.Printf("S == %+v\n", S)
	//赋值
	S.int = 1000 //该匿名函数的赋值  直接赋值就行

}
