package main

import "fmt"

func main() {

	// you can only do this inside a func 你只能在函数内部这样做
	message := "Hello World!"
	//使用短变量声明定义变量 message，并将其赋值为字符串 "Hello World!"
	a, b, c := 1, false, 3
	//同时定义变量 a, b, c，并分别赋值
	d := 4
	//使用短变量声明定义变量 d，并赋值为 4。
	e := true
	//使用短变量声明定义变量 e，并赋值为布尔值 true
	fmt.Println(message, a, b, c, d, e)
}

//这段代码展示了在 Go 语言中如何在函数内部定义和使用多个变量，并使用 fmt.Println 进行打印。
