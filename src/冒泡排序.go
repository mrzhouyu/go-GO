package main

import "fmt"
import "math/rand"
import "time"

func main() {

	//设置随机数种子
	rand.Seed(time.Now().Unix())
	var array [10]int
	n := len(array)
	for i := 0; i < n; i++ {
		array[i] = rand.Intn(20)
	}
	fmt.Println(array)

	for t := 0; t < n-1; t++ {
		for j := 0; j < n-1-t; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]

			}
		}

	}
	fmt.Println(array)

	for m := 0; m < n-1; m++ {
		for k := 0; k < n-1-m; k++ {
			if array[k] < array[k+1] {
				array[k], array[k+1] = array[k+1], array[k]
			}
		}

	}
	fmt.Println(array)

}
