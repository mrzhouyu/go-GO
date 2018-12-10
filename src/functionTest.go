package main

import "fmt"

func MaxMin(num1 int, num2 int) (max, min int) {
	if num1 > num2 {
		max = num1
		min = num2
	} else {
		max = num2
		min = num1
	}
	return

}

func main() {
	// var max, min int

	max, min := MaxMin(10, 20)
	fmt.Printf("max={%d} ,min={%d}\n", max, min)
}
