package main

import "fmt"

func main() {
	var message string
	//声明变量 message，其类型为 string，此时尚未初始化，默认值为 ""（空字符串）
	var a, b, c int
	//声明变量 a, b, c，类型为 int（整数）。此时这些变量的默认值均为 0。
	a = 1
	//变量 a 赋值 1

	message = "Hello World!"
	//变量 message 赋值为字符串 "Hello World!"

	fmt.Println(message, a, b, c)
}
