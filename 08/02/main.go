package main

import "fmt"

func main() {

	a := 43
	//定义与初始化

	fmt.Println(a)  // 43
	fmt.Println(&a) // 0x20818a220
	//打印变量和地址

	var b = &a      //创建指针
	fmt.Println(b)  // 0x20818a220
	fmt.Println(*b) // 43
	//指针的打印

	// b is an int pointer;
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is known as dereferencing
	// the * is an operator in this case
	//fmt.Println(*b) 使用 * 操作符来解引用指针 b，返回 b 指向的值，即 a 的值 43
}

/*指针的好处：使用指针可以在函数间传递变量的引用，避免大块数据的复制，提高效率。
可以通过指针直接修改变量的值，而无需返回新的值。*/
