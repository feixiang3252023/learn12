package main

import "fmt"

func main() {
	x := 42
	//这里首先声明了一个变量 x，并将其初始化为 42
	//使用 := 是 Go 语言中简洁的变量声明方式。
	fmt.Println(x)
	//输出x的值,即42
	{
		fmt.Println(x)
		//fmt.Println(x) 再次输出 x 的值，即 42（因为 x 是在 main 函数的外层作用域中声明的
		y := "The credit belongs with the one who is in the ring."
		//声明新的变量y,并赋值
		fmt.Println(y)
		//输出y的值
	}
	//用 {} 括起来，创建了一个新的作用域
	// fmt.Println(y) // outside scope of y
	//注释，说明如果尝试在 {} 块外部访问 y，将导致编译错误，因为 y 的作用域仅限于创建它的块内部。在 main 函数的外部，无法访问 y。
}
