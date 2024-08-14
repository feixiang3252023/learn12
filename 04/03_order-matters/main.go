package main

import "fmt"

func main() {
	fmt.Println(x)
	//试图打印局部变量 x，但在这一行中，x 还没有被初始化，
	fmt.Println(y)
	//将会打印全局变量 y 的值（42），这是合法的，因为 y 之前已被声明并初始化。
	x := 42
	//x := 42 声明并初始化了局部变量 x，但这个声明是在使用变量的后面，因此不会被执行。
}

var y = 42
