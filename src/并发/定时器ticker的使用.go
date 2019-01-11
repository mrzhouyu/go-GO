package main

import "fmt"
import "time"

func main() {
	//此定时器是循环 定时的作用
	ticker := time.NewTicker(time.Second)
	i := 0
	for {
		<-ticker.C //时间未到时阻塞
		i++
		fmt.Println("i= ", i)
		if i == 5 {
			ticker.Stop()
			break
		}
	}

}
