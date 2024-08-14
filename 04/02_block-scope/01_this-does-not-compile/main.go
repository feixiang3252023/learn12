package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	foo() //调用 foo 函数
}

func foo() {
	// no access to x
	// this does not compile
	fmt.Println(x)
	//尝试打印变量 x 的值
	//由于 x 只在 main 函数的作用域内可用，foo 函数无法访问 x。因此，编译此代码时会出现错误，提示未定义 x。
}
