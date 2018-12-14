package main

import "fmt"

//对比python比较好理解
func fun1() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	s := fun1()
	fmt.Println(s()) //1
	fmt.Println(s()) //4
	fmt.Println(s()) //9

}
