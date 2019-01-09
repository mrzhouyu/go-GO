package main

import "fmt"

//channel可以关闭 关闭后不可以向其写入数据 但是可以读取数据
func main() {
	ch := make(chan int) //无缓存channel
	//新建一个协程
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i //channel写入数据
		}
		close(ch) //不需要写入数据时候 关闭 但是注意关闭完不可以再写入数据 但是可以读取数据
	}()

	//新建一个死循环
	for {
		if num, ok := <-ch; ok == true { //没有关闭返回true
			fmt.Println("num= ", num)
		} else { //关闭管道
			break //跳出循环
		}
	}

	//或者  迭代遍历ch 被关闭则跳出循环
	for num := range ch {
		fmt.Println("num=", num)
	}
}
