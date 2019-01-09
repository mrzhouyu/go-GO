package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string) //chan初始值为nil  这里是 无缓存的channel
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子协程i= ", i)
			time.Sleep(time.Second)

		}
		ch <- "解除阻塞"
	}()

	channel := <-ch //没数据之前阻塞在这里
	fmt.Println("channel= ", channel)

}
