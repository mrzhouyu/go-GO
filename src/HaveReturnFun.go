package main

import "fmt"

//函数名字首字母小写为私有函数，首字母大写为共有变量
//有返回值的函数需要return中断
//只有一个返回值可以省略()
//无返回值得时候可以省略返回信息否则必须有return语句
func myFun1() int {
	return 666
}

//给返回值起一个变量名字，golang推荐写法
func myFun2() (result int) {
	return 666
}

//常用的方法
func myFun3() (result int) {
	result = 666
	return
}

//多个返回
func myFun4() (int, int, int) {
	return 1, 2, 3
}

//多个返回的推荐写法
func myFun5() (a int, b int, c int) {
	a, b, c = 1, 2, 3
	return
}

//第二种
func myFun6() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

func main() {
	a := myFun1()
	fmt.Println(a)
	var b int
	b = myFun2()
	fmt.Println(b)
	c := myFun3()
	fmt.Println(c)

}
