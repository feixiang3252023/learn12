package main

import "fmt"

func main() {

	a := 43
	//变量定义

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a) //打印内存地址
	fmt.Printf("%d \n", &a)
	//使用 Printf 以十进制格式打印内存地址
}
