package main

import "fmt"

func main() {
	//切片的定义方式和属性
	//定义并且初始化
	s1 := []int{1, 2, 3}
	s2 := [...]int{1, 2, 3}
	var s3 []int = []int{1, 2, 3}
	//利用make函数初始化 5为长度，10为容量 类似于python 但是python不用指定容量 不指定容量时候长度和容量相等
	s5 := make([]int, 5, 10)
	s5 = append(s5, 100)
	fmt.Println("make初始化的：", s5)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Printf("属性： 长度len: %d   ,容量cap:%d   \n", len(s3), cap(s3))
	//数组和切片
	array := [5]int{1, 2, 3, 4, 5}
	s4 := array[1:4:5]
	fmt.Println(s4)

}
