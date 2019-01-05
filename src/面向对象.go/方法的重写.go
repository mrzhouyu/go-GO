package main

import "fmt"

//自定义的类型 这里是结构体
type Person struct {
	name string
	age  int
	sex  byte
}

//属于Person的一个方法（对象接收者）
func (r *Person) PrintInfo() {
	fmt.Printf("name=%s , age=%d, sex=%c \n", r.name, r.age, r.sex)
}

//定义一个子继承 继承了Person的方法和字段
type Student struct {
	Person //匿名字段
	id     int
	addre  string
}

//重写原本属于父级的方法 方法同名
func (r *Student) PrintInfo() {
	fmt.Println(r)
}
func main() {

	S := Student{Person{"jimi", 18, 'w'}, 1, "earth"}
	S.PrintInfo()        //这里调用的是重写之后的方法  就近原则
	S.Person.PrintInfo() //显式调用 属于父级的方法
}
