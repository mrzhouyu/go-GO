package main

//结构体作为参数传递是拷贝 即使在函数内部被改变  也不会改变主函数的值  如果想改变 可以传递指针地址
import "fmt"

type Man struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//全部初始化
	var student Man = Man{1, "chou", 'm', 10, "china"}
	teacher := Man{name: "chou", sex: 'm'}
	//指针结构体
	var Student *Man = &Man{1, "chou", 'm', 10, "china"}

	fmt.Println("student: ", student)
	fmt.Println("teacher: ", teacher)
	fmt.Println("*Student: ", *Student)
	//声明并调用
	var S Man
	S.id = 1
	S.name = "chou"
	S.age = 10
	S.sex = 'm'
	S.addr = "china"
	//指针
	var p *Man
	p = &S
	//(*p).id和S.id完全等价
	(*p).id = 1
	//结构体的比较只能 ==  或者!=比较 不能< 或者>

	//

}
