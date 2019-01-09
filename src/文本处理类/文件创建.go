package main

import "fmt"
import "os"
import "io"
import "bufio"

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

//读取整个

func ReadFile(path string) {
	f, err1 := os.Open(path)
	if err1 != nil {
		fmt.Println("err= ", err1)
		return
	}
	buf := make([]byte, 2048) //定义一个2k大小的切片
	n, err2 := f.Read(buf)    //读取的文件给buf并且返回buf的长度
	if err2 != nil && err2 != io.EOF {
		fmt.Println("err2= ", err2)
		return
	}
	fmt.Println("buf = \n", string(buf[:n])) //只要内容的长度 且字符串化切片
	defer f.Close()
}

//按行读取
func ReadFileLine(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err= ", err)
		return
	}
	defer f.Close()
	//新建一个缓冲区 先把内容放在缓冲区
	r := bufio.NewReader(f)
	for {
		buf, err1 := r.ReadBytes('\n') //遇到换行结束读取读取内包括\n 注意这里用的是单引号 表示字符与python有区别 返回一个切片
		if err1 != nil {
			if err1 == io.EOF { //匹配文件结束符 表示文件整个读取完毕
				break
			}
			fmt.Println("err1= ", err1)
		}
		fmt.Printf("buf = %s\n", string(buf))

	}
}


func main() {
	path := "demo.txt"
	// WriteFile(path)
	//ReadFile(path)
	ReadFileLine(path)

}
