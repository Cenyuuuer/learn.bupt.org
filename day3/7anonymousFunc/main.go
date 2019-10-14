package main

import (
	"fmt"
)

func add(a int, b int) int {
	// sum := func(a int, b int) int {
	// 	return a + b
	// }(a, b)
	sum := func(a int, b int) int {
		return a + b
	}
	return sum(a, b)
}
func main() {
	sum := add(1, 2)
	fmt.Println(sum)
}
