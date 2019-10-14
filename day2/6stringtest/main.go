package main

import (
	"fmt"
)

func reverse(a string) (result string) {
	strlen := len(a)
	for i := 0; i < len(a); i++ {
		result = result + fmt.Sprintf("%c", a[strlen-i-1])
	}
	return
}
func main() {
	var str1 = "hello  "
	str2 := "world"
	str3 := str1 + str2
	fmt.Println(str3)
	fmt.Sprintf("%s %s\n", str1, str2)
	n := len(str1)
	fmt.Println(n)

	//切片
	substr := str3[:5]
	fmt.Println(substr)
	substr = str3[6:]
	fmt.Println(substr)
	//翻转
	str4 := reverse(str3)
	fmt.Println(str4)
}
