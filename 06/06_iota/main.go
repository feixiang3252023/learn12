package main

import "fmt"

const (
	_ = iota // 0
	//这一行是占位符，它的值为 0，但因为 _ 表示我们不需要这个值，所以它不会被使用。
	b = iota * 10 // 1 * 10
	c = iota * 10 // 2 * 10
	d = iota * 10 // 3 * 10
	e
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}
