package main

import (
	"fmt"
)

func swap(a int, b int) {
	a, b = b, a
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return
}
func swap1(a *int, b *int) {
	a, b = b, a
	return
}
func main() {
	a := 100
	b := 200
	// swap(a, b)
	swap1(&a, &b)
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
}
