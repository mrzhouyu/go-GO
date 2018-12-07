package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
	)
	fmt.Println("a = ", a, "b =", b, "c=", c)

	t := 1 + 2i
	fmt.Println("t=", t)

	fmt.Println("real = ", real(t), "imag=", imag(t))
}
