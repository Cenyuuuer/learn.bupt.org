package main

import (
	"fmt"
	"time"
)

const (
	male   = 1
	female = 2
)

func main() {
	second := time.Now().Unix()
	fmt.Println(second)
	if second%female == 0 {
		fmt.Println("female")
	} else {
		fmt.Println("male")
	}

}
