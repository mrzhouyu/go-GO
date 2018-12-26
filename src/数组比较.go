package main

import "fmt"

func main() {

	//只支持比较 ==  或 != 比较的是每个元素是否一样  如果两个数组比较数组类型要是一样
	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}
	fmt.Println("a==b", a == b)
	fmt.Println("a!=b", a != b)

}
