//调用	"io/ioutil" 库直接先把 文件读入内存 再使用split分割
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	count := make(map[string]int)
	for _, filename := range os.Args[1:] {
		date, err := ioutil.ReadFile(filename) //返回的是 byte slice类型 需要转化为字符串 err是错误参数 若无措 err == nil
		if err != nil {
			fmt.Fprintf(os.Stderr, "readFile3 :%v\n", err) //加一个错误处理
			continue
		}
		for _, line := range strings.Split(string(date), "\n") {
			count[line]++
		}
	}

	for line, n := range count {
		if n > 0 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
