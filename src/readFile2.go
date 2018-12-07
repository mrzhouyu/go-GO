//读取文件
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var fileName string = os.Args[0]

func main() {
	dic := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		count(os.Stdin, dic)
	} else {
		for _, arg := range files {
			fmt.Printf("i have got file %s\n", arg)
			f, err := os.Open(arg) //返回两个参数 一个是被打开的文件 一个是内置的错误类型
			if err != nil {
				fmt.Println("err is be found!")
				fmt.Fprint(os.Stderr, "readfile2 : %v \n", err)
				continue
			}
			count(f, dic)
			f.Close()
		}
	}
	for line, n := range dic {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

//
func count(f *os.File, dic map[string]int) { //*os.File 结构体
	input := bufio.NewScanner(f) //创建bufio.NewScanner类型的变量
	for input.Scan() {           //然后被Scan读取一行  每次读取一行且去除末尾的 \n符号
		repatOutName(input.Text(), dic)
		dic[input.Text()]++ // Scan  读完一行返回true  否则false     用.Text()方法获取 值

	}
}

//加一个判断 出现重复的行时候打印 文件名称

func repatOutName(s string, d map[string]int) {
	for line, _ := range d {
		if s == line {
			slice := strings.Split(fileName, "\\") //splite后返回的是一个切片对象
			var l int = len(slice) - 1             //获取总长度 -1 最后一个元素的位置

			fmt.Printf(" ‘%s’ is repeat ,print fileName: %s\n", s, slice[l])
			break
		}
	}
}
