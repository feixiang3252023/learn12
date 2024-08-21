package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	// 打印 z 的地址（在内存中的地址）
	fmt.Println(&z) // address in func zero
	z = 0
	// 修改 z 的值，此处不会影响函数外的 x
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	//打印 x 的地址（在内存中的地址）
	fmt.Println(&x) // address in main
	zero(x)         //调用 zero 函数，传递 x 的值
	fmt.Println(x)  // x is still 5
}
