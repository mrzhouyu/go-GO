package main

import "fmt"
import "os"

func WriteFile(path string) {
	//打开文件，新建文件
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("create err= ", err)
		return
	}
	//函数调用结束前关闭
	defer f.Close()

	var str string
	for i := 0; i < 10; i++ {
		//生成字符串
		str = fmt.Sprintf("n=%d\n", i+1)
		//n代表的是字符长度
		n, err := f.WriteString(str)
		if err != nil {
			fmt.Println("写入出错")
			return
		}
		fmt.Println("n=", n)

	}
}

func main() {
	path := "demo.txt"
	WriteFile(path)
}
