package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	second := time.Now().Second()
	fmt.Printf("%02d %02d", now.Year(), second)
}
