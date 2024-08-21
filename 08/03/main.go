package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43 输出变量 a 的值，结果是 43
	fmt.Println(&a) // 0x20818a220 输出变量 a 的内存地址

	var b = &a      //声明一个指针 b，指向 a 的地址
	fmt.Println(b)  // 0x20818a220 输出 b 的值，即 a 的内存地址
	fmt.Println(*b) // 43 解引用 b，输出 43，即 a 的当前值

	*b = 42 // b says, "The value at this address, change it to 42"
	//修改 b 指向的内存地址中的值
	fmt.Println(a) // 42
	//输出新的 a 的值

	// this is useful 指针的用处
	// we can pass a memory address instead of a bunch of values (we can pass a reference)
	//我们可以传递一个内存地址，而不是多个值（可以通过引用进行传递）
	// and then we can still change the value of whatever is stored at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
	// we only have to pass around addresses

	// everything is PASS BY VALUE in go, btw
	// when we pass a memory address, we are passing a value
}

//Go 语言中所有的函数参数都是通过值传递，即使是指针，传递的是指针的值（内存地址），而不是直接的内存地址本身。
