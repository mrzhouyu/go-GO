package main

import (
	"fmt"
	"regexp"
)

func main() {
	s := "bacabbcbabbabcadcaddccajcaoc"
	//解释规则 解析正则表达式 如果成功返回解释器
	part := regexp.MustCompile(`a.c`) //类似于python
	if part == nil {                  //失败返回nil
		fmt.Println("匹配错误")
		return //如果匹配不到则结束程序
	}
	//成功根据规则提前关键信息 -1表示返回所有查找到的内容
	r := part.FindAllStringSubmatch(s, -1)
	fmt.Println(r)

}
