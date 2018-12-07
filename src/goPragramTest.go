package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	S, SEP := "", ""
	fmt.Println(strings.Join(os.Args[1:], "--"))
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println("first: " + s)

	//这里用_表示这个量可以不被使用，改为其他的变量必须被使用 如 x
	for _, arg := range os.Args[1:] {
		// fmt.Println(x)
		S += SEP + arg
		SEP = " "
	}
	fmt.Println("second :" + S)
}
