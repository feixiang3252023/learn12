package main

import "fmt"

const p = "death & taxes"

//定义了一个包级别的常量p，其值为字符串"death & taxes"
//常量在Go中是一个在程序运行时不变的值。

func main() {

	const q = 42
	//在main函数内部，这行代码定义了另一个常量q，其值为整数42。此常量的作用域仅限于main函数。

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)
	//打印出常量p和q的值到控制台
}

// a CONSTANT is a simple unchanging value
