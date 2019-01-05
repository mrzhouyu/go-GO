package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 string = "hello world      "
	fmt.Println(strings.Fields(str1)) //

	fmt.Println(strings.Contains(str1, "wor")) //判断字符串是否包含"wor"是就返回true

	fmt.Println(strings.Index(str1, "wor")) //找到wor在字符串的位置        从零开始 不存在返回-1

	var sli []string = []string{"one", "two", "three"}

	fmt.Println(strings.Join(sli, "_")) //根据某个字符连接切片

	fmt.Println(strings.Repeat("ha", 3)) //重复ha三次返回

	fmt.Println(strings.Split(str1, " ")) //拆分字符串

	fmt.Println(strings.Trim(str1, " ")) //去除两端指定字符

	fmt.Println(strings.Fields(str1)) //去掉空格 返回一个切片

}
