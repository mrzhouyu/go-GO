package main

import "fmt"

func func1(a [5]int) {
	a[0] = 100
	fmt.Println("传递数组元素改变a[0]: ", a)
}

func func2(p *[5]int) {
	(*p)[0] = 100
	fmt.Println("加括号：", *p)
}
func main() {

	var a [5]int = [5]int{1, 2, 3, 4, 5}

	//传递参数 相当于拷贝一份 不改变主函数的值
	func1(a)
	fmt.Println("没有改变主函数的值吧： ", a)
	//传递地址，会改变主函数数组的值
	func2(&a)
	fmt.Println("传递地址后值有没有改变: ", a)

}
