package main

import "fmt"

//map存储是无序的 底层是hash
//map定义 1 var name map[key_type]value_type
var hi map[string]int

//map 使用之前需要进行初始化 才能被调用生效 否则报错
//初始化函数
func init() {
	hi = make(map[string]int)
}

func main() {
	//map定义 2 name := make(map[key_type]value_type)
	// hello := make(map[string]int)

	hi["a"] = 100
	fmt.Println("Map: ", hi)
	fmt.Println("value:", hi["a"])
}
