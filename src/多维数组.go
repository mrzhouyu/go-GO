package main

import "fmt"

func main() {

	var a [3][4]int //定义
	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
		}
	}
	fmt.Println(a)

	//初始化 没有初始化的值为0
	b := [3][4]int{{1, 2, 3, 4}, {}}
	fmt.Println(b)

	c := [3][4]int{1: {1: 2, 3: 4}, 2: {3: 9}}
	fmt.Println(c)

}
