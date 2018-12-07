//从标准输入里读取 即：从命令行读取
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	fmt.Println("Please input some Worlds: ")
	for input.Scan() {
		counts[input.Text()]++
		if input.Text() == "over" {
			break
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
