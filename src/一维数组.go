package main

import "fmt"

func main() {
	var array [20]int                          //先定义 后面再赋值
	var array1 [5]int = [20]int{1, 2, 3, 4, 5} //全部初始化初始化 边定义边赋值
	array2 := [5]int{1, 2, 3, 4, 5}            //初始化自动推到
	array3 := [5]int{1, 2}                     //部分初始化 没有赋值的自动为0
	array4 := [5]int{1: 1, 2: 2}               //指定下标初始化 没有初始化的为0

	for i := 0; i < len(array); i++ {
		array[i] = i + 1
		// fmt.Printf("array[%d]=%d\n", i, array[i])
	}

	for index, value := range array { //数组的遍历，返回两个值 一个索引和对应的值，如果不需要索引可以用 "_"赋值
		fmt.Printf("index = %d, value = %d \n", index, value)
	}

}
