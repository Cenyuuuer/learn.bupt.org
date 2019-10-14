package main

import (
	"fmt"
)

func main() {
	var a string = "hello 	world"
	var b string = `hello
	world`
	var c byte = 'c'
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%c\n", c)
	fmt.Println(c)
}
