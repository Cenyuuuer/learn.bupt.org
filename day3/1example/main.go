package main

import (
	"fmt"
	"strconv"
	"strings"
)

func urlProcess(url string) string {
	result := strings.HasPrefix(url, "http://")
	if !result {
		url = fmt.Sprintf("http://%s", url)

	}
	return url
}

func pathProcess(path string) string {
	result := strings.HasSuffix(path, "http://")
	if !result {
		path = fmt.Sprintf("%s/", path)

	}
	return path
}

func main() {
	var (
		url  string
		path string
	)
	fmt.Scanf("%s%s", &url, &path)
	url = urlProcess(url)
	path = pathProcess(path)
	fmt.Println(url)
	fmt.Println(path)

	var str string = "a b c"
	sliceStr := strings.Fields(str) //还有split
	for _, value := range sliceStr {
		fmt.Println(value)
	}

	//吧字符串转换成整数
	str1 := strconv.Itoa(100)
	fmt.Println(str1)
	//把整数转换成字符串
	int1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("can not convert")
	}
	fmt.Println(int1)
}
