package main

import "fmt"
import "math/rand"
import "time"

func main() {

	rand.Seed(100) //如果seed是固定的 那么每次产生的随机数也是一样的
	for i := 0; i < 3; i++ {
		fmt.Println("随机数是: ", rand.Intn(100)) //100 以内
	}

	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 3; j++ {
		fmt.Println("时间随机数产生: ", rand.Intn(100))
	}
}
