package main

import "fmt"

var x = 42

//声明变量 x 并初始化为整数 42。这个变量在整个包中都是可见的。

func main() {
	fmt.Println(x)
	foo() //调用 foo 函数
}

func foo() { //定义 foo 函数
	fmt.Println(x)
}
