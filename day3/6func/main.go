package main

import (
	"fmt"
)

//其实arg是一个切片，可以通过arg[index]来访问参数，通过len(arg)来判断参数个数
func add1(arg ...int) int { //0个或多个参数
	size := len(arg)
	var sum int
	for i := 0; i < size; i++ {
		sum += arg[i]
	}
	return sum
}

// func add2(a int,arg...int)int{ //1个或多个参数

// }
// func add3(a int,b int ,arg..int)int{//2个或多个参数

// }
func main() {
	sum := add1(1, 2)
	fmt.Println(sum)
}
