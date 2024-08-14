package main

import "fmt"

func wrapper() func() int { //定义一个名为 wrapper 的函数，它返回一个匿名函数
	x := 0 //在 wrapper 中，局部变量 x 被初始化为 0
	return func() int {
		x++
		return x
	}
	//函数返回的匿名函数也访问这个局部变量 x，并且每次调用它时都会将 x 增加 1
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

//在 main 函数中调用 wrapper 函数，并将返回的匿名函数赋值给 increment。

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
