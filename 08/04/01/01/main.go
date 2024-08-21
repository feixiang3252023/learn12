package main

import "fmt"

func zero(z int) {
	z = 0 // 这里修改的是 z 的副本，不会影响 x
}

func main() {
	x := 5         // 初始化变量 x，其值为 5
	zero(x)        // 调用 zero 函数，传递 x 的值
	fmt.Println(x) // x is still 5
}
