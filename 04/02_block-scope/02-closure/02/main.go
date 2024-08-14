package main

import "fmt"

var x = 0

//声明一个全局变量 x，并初始化为 0。这个变量可以在包中的任意位置访问。

func increment() int { //定义一个名为 increment 的函数
	x++ //将全局变量 x 的值加 1。
	return x
}

func main() {
	fmt.Println(increment())
	fmt.Println(increment())
}

//定义 main 函数，这是程序的执行入口。它调用 increment 函数两次，并打印返回值。

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
