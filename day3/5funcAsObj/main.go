package main

import (
	"fmt"
)

func add(a int, b int) int {
	return a + b
}

type add_func func(a int, b int) int //定义一个类型，类型的名字叫做add_func

func oprater(op add_func, a int, b int) int {
	return op(a, b)
}
func main() {
	var (
		a int = 1
		b int = 2
	)
	c := add
	d := c(a, b)
	fmt.Println(d)

	sum := oprater(c, 100, 200)
	fmt.Println(sum)
}
