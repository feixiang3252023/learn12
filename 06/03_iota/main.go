package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
	//iota 是 Go 语言中的一个预定义标识符，用于生成一组常量的递增数字。每当你在常量声明中使用它时，它的值会自动递增
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
