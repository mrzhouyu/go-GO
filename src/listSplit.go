package main

import "fmt"
package main
/*
特点：
内存连续的数据集合 var arr [5]int --> arr[0] ...arr[4]
*/

//定义数组
//定义一个 长度为2 的数组
var arr1 [2]int

func main() {
	arr1[0] = 1
	arr1[1] = 2

	fmt.Println("输出：", arr1)
}
