package main

import "fmt"

func zero(z *int) {
	*z = 0 // 通过指针修改 z 指向的值
}

func main() {
	x := 5         // 定义一个整型变量 x 并赋值为 5
	zero(&x)       // 将 x 的地址传递给 zero 函数
	fmt.Println(x) // x is 0
}
