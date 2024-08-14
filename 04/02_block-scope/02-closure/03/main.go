package main

import "fmt"

func main() {
	x := 0
	//声明一个局部变量 x，并初始化为 0
	increment := func() int { //定义一个匿名函数并将其赋值给变量 increment
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable
*/
/*闭包：这里的匿名函数（increment）形成了闭包，可以访问其外部的变量 x，尽管 x 是在 main 函数中定义的。这限制了 x 的作用域，仅在 increment 函数中使用。
匿名函数：没有名称的函数，通过赋值给变量以便后续调用。
函数表达式：将一个函数赋值给一个变量，使其可以像普通变量一样使用。*/
