package main

import "fmt"

func MyFun(a, b int) { //普通有参数无返回
	fmt.Println(a)
	fmt.Println(b)

}

func Arg1(args ...int) { //不定参数无返回值
	fmt.Println(len(args))
	for i := 0; i < len(args); i++ {
		fmt.Printf("%d\n", args[i])
	}
	for i, value := range args {
		fmt.Printf("args[%d] = %d\n", i, value)
	}

}

func Arg2(x int, args ...int) { //不定参数和参数都有的话 ，餐数要放在前面，不定参数凡在后面
	fmt.Println(len(args))
}

func Arg3(tmp ...int) { //不定参数传给不定参数函数传递
	Arg2(tmp...)
	Arg2(tmp[2:]) //传递最后两个参数到下一个函数

}

func main() {
	// MyFun(1, 2)
	Arg1()
	Arg1(1, 2, 3, 4)
	//Arg2()  这样是错误的 必须要有参数
	Arg2(1, 2, 3, 4, 5)
}
