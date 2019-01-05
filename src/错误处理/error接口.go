package main

import "fmt"
import "errors"

//定义一个函数 error的应用
func Div(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0。")
	} else {
		result = a / b
	}
	return
}

func main() {
	var err error = fmt.Errorf("%s", "这样可以定义一个error不推荐")
	err1 := fmt.Errorf("%s", "这是定义一个error的第二种方式推荐")
	err2 := errors.New("这是定义一个error的第三种方式，导入errors来实现")
	fmt.Printf("err=%v\nerr1=%v\nerr2=%v\n", err, err1, err2)

	//以下为error的具体应用
	r, err := Div(1, 2)
	if err != nil {
		fmt.Println("err=", err)
	} else {
		fmt.Println("r=", r)
	}
}
