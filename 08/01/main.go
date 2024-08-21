package main

import (
	"fmt"
)

func main() {

	a := 43
	//声明了一个整型变量 a，并赋值为 43

	fmt.Println(a)
	fmt.Println(&a)
	//打印变量和指针

	var b = &a
	//创建指针变量

	fmt.Println(b)
	//打印指针变量

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
	//总结了指针的基本概念。b 是一个指向整型的指针，*int 表示它指向一个整数的地址
}
