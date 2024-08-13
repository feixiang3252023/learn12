package main

import "fmt"

func main() {

	var message = "Hello World!"
	//使用 var 关键字声明一个变量 message，并将其初始化为字符串 "Hello World!"
	var a, b, c = 1, false, 3
	//同时声明变量 a, b, c，并分别赋值。a 是整数 1，b 是布尔值 false，c 是整数 3

	fmt.Println(message, a, b, c)
}
