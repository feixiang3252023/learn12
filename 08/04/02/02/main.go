package main

import "fmt"

func zero(z *int) {
	fmt.Println(z) // 打印指针 z 的地址
	*z = 0         // 通过指针修改 z 指向的值，将其设置为 0
}

func main() {
	x := 5 // 定义一个整型变量 x，并赋值为 5
	fmt.Println(&x)
	zero(&x)       // 将 x 的地址传递给 zero 函数
	fmt.Println(x) // x is 0
}
