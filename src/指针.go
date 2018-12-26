package main

import "fmt"

func swap1(a, b int) {
	a, b = b, a
	fmt.Printf("a=%v,b=%v\n", a, b)
}

func swap2(p1, p2 *int) {
	*p1, *p2 = *p2, *p1
	fmt.Printf("*p1=%v,*p2=%v\n", *p1, *p2)
}

func main() {
	var p *float32
	p = new(float32)
	*p = 3.1415
	fmt.Printf("&p = %v\n", &p)
	fmt.Printf("*p = %v\n", *p)
	a, b := 10, 20

	swap1(a, b) //用参数传递
	fmt.Printf("a=%v,b=%v\n", a, b)

	swap2(&a, &b) //用指针传递
	fmt.Printf("a=%v,b=%v\n", a, b)

}
