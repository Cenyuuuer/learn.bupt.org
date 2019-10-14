package main

import "fmt"

func main() {

	var a int = 5

	switch a {
	case 0, 3: //可以多个条件
		fmt.Println("h")
		fmt.Println("e")
		fallthrough //会直接加入下一个case
	case 2:
		fmt.Println("a")
	default:
		fmt.Println("default") //当没有符合的case的时候进入default

	}

	switch {
	case a < 2:
		fmt.Println("hhh")
	case a >= 2:
		fmt.Println("aaa")
	}
}
